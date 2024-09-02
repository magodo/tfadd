package internal

import (
	"fmt"
	"sort"
	"strings"

	"github.com/magodo/tfadd/schema"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"

	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	tfpluginschema "github.com/magodo/tfpluginschema/schema"
	"github.com/zclconf/go-cty/cty"
)

func TuneTpl(sch schema.Schema, tpl []byte, option *TuneOption) ([]byte, error) {
	t, err := newTrimmer(tpl, option)
	if err != nil {
		return nil, err
	}
	if err := t.tuneBlock(t.Body, sch.Block, Addr{}); err != nil {
		return nil, err
	}
	return t.Bytes(), nil
}

type trimmer struct {
	File   *hclwrite.File
	Body   *hclwrite.Body
	Option TuneOption
}

func newTrimmer(tpl []byte, option *TuneOption) (*trimmer, error) {
	f, diag := hclwrite.ParseConfig(tpl, "", hcl.InitialPos)
	if diag.HasErrors() {
		return nil, fmt.Errorf("parsing the template: %s", diag.Error())
	}
	if n := len(f.Body().Blocks()); n != 1 {
		return nil, fmt.Errorf("invalid template: expect one top level block, got=%d", n)
	}
	rb := f.Body().Blocks()[0].Body()

	rb.RemoveAttribute("id")
	rb.RemoveBlock(rb.FirstMatchingBlock("timeouts", nil))

	var opt TuneOption
	if option != nil {
		opt = *option
	}
	if opt.OCToKeep == nil {
		opt.OCToKeep = map[string]bool{}
	}

	return &trimmer{
		File:   f,
		Body:   rb,
		Option: opt,
	}, nil
}

func (t trimmer) Bytes() []byte {
	return t.File.Bytes()
}

func (t trimmer) removeAttribute(body *hclwrite.Body, addr Addr, name string) {
	if t.Option.OCToKeep[addr.String()] {
		return
	}
	body.RemoveAttribute(name)
}

func (t trimmer) removeBlock(body *hclwrite.Body, addr Addr, blk *hclwrite.Block) {
	if t.Option.OCToKeep[addr.String()] {
		return
	}
	body.RemoveBlock(blk)
}

func (t trimmer) tuneAttributes(parentAddr Addr, rb *hclwrite.Body, attrSchs tfpluginschema.SchemaAttributes) error {
	schMap := attrSchs.Map()
	for attrName, attrVal := range rb.Attributes() {
		addr := parentAddr.AppendAttributeStep(attrName)

		sch, ok := schMap[attrName]
		if !ok {
			// This might because the provider under used is a newer one than the version where we ingest the schema information.
			// This might happen when the user has a newer version provider installed in its local fs, and has set the "dev_overrides" for that provider.
			// We simply remove that attribute from the config.
			t.removeAttribute(rb, addr, attrName)
			continue
		}

		// Always remove C only attribute
		if sch.Computed && !sch.Optional {
			t.removeAttribute(rb, addr, attrName)
			continue
		}

		// Removing O+C attribute
		if t.Option.RemoveOC {
			if sch.Computed && sch.Optional {
				// Removing O+C attributes as long as they meet the property constraints
				if len(sch.ExactlyOneOf) != 0 {
					// For O+C attribute that has "ExactlyOneOf" constraint, keeps the first one in alphabetic order.
					l := make([]string, len(sch.ExactlyOneOf))
					copy(l, sch.ExactlyOneOf)
					sort.Strings(l)

					if l[0] != addr.String() {
						t.removeAttribute(rb, addr, attrName)
						continue
					}
				} else if len(sch.AtLeastOneOf) == 0 {
					// For O+C attribute that has "AtLeastOneOf" constraint, keep it
					t.removeAttribute(rb, addr, attrName)
					continue
				}
			}
		}

		// Removing Optional "zero" valued attribute
		if t.Option.RemoveOZAttribute {
			if sch.NestedType == nil {
				ok, err := t.attributeIsDefaultOrZeroValue(attrName, attrVal, sch)
				if err != nil {
					return addr.NewErrorf("checking attribute value is default or zero: %v", err)
				}
				if ok {
					t.removeAttribute(rb, addr, attrName)
					continue
				}
			}
		}

		// TODO: Attributes that are kept (either kept O+C, O or R attributes), continue trim the nested objects, or trim the attribute by value.
		if sch.NestedType != nil {
		}
	}
	return nil
}

func (t trimmer) tuneBlock(rb *hclwrite.Body, sch *tfpluginschema.SchemaBlock, parentAddr Addr) error {
	if sch == nil {
		return nil
	}

	if err := t.tuneAttributes(parentAddr, rb, sch.Attributes); err != nil {
		return parentAddr.NewErrorf("tunning attributes: %v", err)
	}

	for _, blk := range rb.Blocks() {
		sch := sch.BlockTypes.Map()[blk.Type()]
		addr := parentAddr.AppendBlockStep(blk.Type())

		// Always remove C only attribute
		if (sch.Computed != nil && *sch.Computed) && !(sch.Optional != nil && *sch.Optional) {
			t.removeBlock(rb, addr, blk)
			continue
		}

		if t.Option.RemoveOC {
			if sch.Computed != nil && *sch.Computed && sch.Optional != nil && *sch.Optional {
				// Removing O+C blocks as long as they meet the property constraints
				if len(sch.ExactlyOneOf) != 0 {
					// For O+C blocks that has "ExactlyOneOf" constraint, keeps the first one in alphabetic order.
					l := make([]string, len(sch.ExactlyOneOf))
					copy(l, sch.ExactlyOneOf)
					sort.Strings(l)

					if l[0] != addr.String() {
						t.removeBlock(rb, addr, blk)
						continue
					}
				} else if len(sch.AtLeastOneOf) == 0 {
					// For O+C attribute that has "AtLeastOneOf" constraint, keep it
					t.removeBlock(rb, addr, blk)
					continue
				}
			}
		}

		if err := t.tuneBlock(blk.Body(), sch.Block, addr); err != nil {
			return addr.NewErrorf("tunning blocks: %v", err)
		}
	}
	return nil
}

// attributeIsDefaultOrZeroValue returns if the attribute is null, or equals to either its default value (defined in schema), or (default value undefined) zero value.
func (t trimmer) attributeIsDefaultOrZeroValue(attrName string, attrVal *hclwrite.Attribute, attrSch *tfpluginschema.SchemaAttribute) (bool, error) {
	if attrSch.NestedType != nil {
		panic("attributes of nested object are not supported")
	}
	aval, err := t.attrValue(attrName, attrVal)
	if err != nil {
		return false, err
	}
	if aval.IsNull() {
		return true, nil
	}

	// Non null attribute, continue checking whether it equals to the default value.
	var dval cty.Value

	if attrSch.Type != nil {
		switch *attrSch.Type {
		case cty.Number:
			dval = cty.Zero
		case cty.Bool:
			dval = cty.False
		case cty.String:
			dval = cty.StringVal("")
		default:
			if attrSch.Type.IsListType() {
				dval = cty.ListValEmpty(attrSch.Type.ElementType())
				if len(aval.AsValueSlice()) == 0 {
					aval = dval
				} else {
					aval = cty.ListVal(aval.AsValueSlice())
				}
				break
			}
			if attrSch.Type.IsSetType() {
				dval = cty.SetValEmpty(attrSch.Type.ElementType())
				if len(aval.AsValueSlice()) == 0 {
					aval = dval
				} else {
					aval = cty.SetVal(aval.AsValueSlice())
				}
				break
			}
			if attrSch.Type.IsMapType() {
				dval = cty.MapValEmpty(attrSch.Type.ElementType())
				if len(aval.AsValueMap()) == 0 {
					aval = dval
				} else {
					aval = cty.MapVal(aval.AsValueMap())
				}
				break
			}
		}
	}

	if attrSch.Default != nil && attrSch.Type != nil {
		var err error
		dval, err = gocty.ToCtyValue(attrSch.Default, *attrSch.Type)
		if err != nil {
			return false, fmt.Errorf("converting cty value %v to Go: %v", attrSch.Default, err)
		}
	}

	return aval.Equals(dval).True(), nil
}

func (t trimmer) attrValue(attrName string, attr *hclwrite.Attribute) (cty.Value, error) {
	attrExpr, diags := hclwrite.ParseConfig(attr.BuildTokens(nil).Bytes(), "generate_attr", hcl.InitialPos)
	if diags.HasErrors() {
		return cty.Zero, fmt.Errorf(`building attribute %q attribute: %s`, attrName, diags.Error())
	}
	attrValLit := attrExpr.Body().GetAttribute(attrName).Expr().BuildTokens(nil).Bytes()
	dexpr, diags := hclsyntax.ParseExpression(attrValLit, "", hcl.InitialPos)
	if diags.HasErrors() {
		return cty.Zero, fmt.Errorf(`parsing HCL expression %q: %s`, string(attrValLit), diags.Error())
	}
	aval, diags := dexpr.Value(&hcl.EvalContext{Functions: map[string]function.Function{
		"jsonencode": stdlib.JSONEncodeFunc,
	}})
	if diags.HasErrors() {
		return cty.Zero, fmt.Errorf(`evaluating value of HCL expression %q: %s`, string(attrValLit), diags.Error())
	}
	return aval, nil
}

type Addr struct {
	segs    []string
	inBlock bool
}

func (addr Addr) NewErrorf(msg string, a ...any) error {
	return fmt.Errorf(fmt.Sprintf("(%s) %s", strings.Join(addr.segs, "."), msg), a...)
}

func (addr Addr) AppendAttributeStep(step string) Addr {
	return addr.appendStep(step, true)
}

func (addr Addr) AppendBlockStep(step string) Addr {
	return addr.appendStep(step, true)
}

func (addr Addr) appendStep(step string, inBlock bool) Addr {
	nsegs := append([]string{}, addr.segs...)
	if addr.inBlock {
		nsegs = append(nsegs, "0")
	}
	nsegs = append(nsegs, step)
	return Addr{segs: nsegs, inBlock: inBlock}
}

func (addr Addr) String() string {
	return strings.Join(addr.segs, ".")
}
