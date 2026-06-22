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

		// Optional only primitive types
		if sch.Optional && sch.NestedType == nil {
			// Removing Optional attribute whose value equals to its zero value
			if sch.Default == nil && t.Option.RemoveOZeroAttribute {
				aval, err := t.attrCty(attrName, attrVal, sch)
				if err != nil {
					return addr.NewErrorf("parsing attribute value: %v", err)
				}
				if attributeIsZeroValue(aval, sch) {
					t.removeAttribute(rb, addr, attrName)
					continue
				}
			}

			// Removing Optional attribute whose value equals the schema-defined default.
			if t.Option.RemoveODefaultAttribute {
				aval, err := t.attrCty(attrName, attrVal, sch)
				if err != nil {
					return addr.NewErrorf("parsing attribute value: %v", err)
				}
				isDefault, err := attributeIsDefaultValue(aval, sch)
				if err != nil {
					return addr.NewErrorf("checking attribute equals schema default: %v", err)
				}
				if isDefault {
					t.removeAttribute(rb, addr, attrName)
					continue
				}
			}
		}

		// TODO(Proto6 only): Attributes that are kept (either kept O+C, O or R attributes), continue trim the nested objects, or trim the attribute by value.
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

// attrCty parses the attribute value and, for collection types, normalizes
// HCL tuple/object literals (e.g. `[]`, `{}`) into properly typed List/Set/Map
// values so that they can be compared with their schema-defined defaults or
// type zero values.
func (t trimmer) attrCty(attrName string, attrVal *hclwrite.Attribute, attrSch *tfpluginschema.SchemaAttribute) (cty.Value, error) {
	if attrSch.NestedType != nil {
		panic("attributes of nested object are not supported")
	}
	aval, err := t.attrValue(attrName, attrVal)
	if err != nil {
		return cty.NilVal, err
	}
	if aval.IsNull() || attrSch.Type == nil {
		return aval, nil
	}
	ty := *attrSch.Type
	switch {
	case ty.IsListType():
		if len(aval.AsValueSlice()) == 0 {
			return cty.ListValEmpty(ty.ElementType()), nil
		}
		return cty.ListVal(aval.AsValueSlice()), nil
	case ty.IsSetType():
		if len(aval.AsValueSlice()) == 0 {
			return cty.SetValEmpty(ty.ElementType()), nil
		}
		return cty.SetVal(aval.AsValueSlice()), nil
	case ty.IsMapType():
		if len(aval.AsValueMap()) == 0 {
			return cty.MapValEmpty(ty.ElementType()), nil
		}
		return cty.MapVal(aval.AsValueMap()), nil
	}
	return aval, nil
}

// attributeIsZeroValue reports whether the given (parsed and normalized)
// attribute value is the "zero" value for an Optional attribute.
// Note that null is not a zero value.
func attributeIsZeroValue(aval cty.Value, attrSch *tfpluginschema.SchemaAttribute) bool {
	var dval cty.Value
	if attrSch.Type == nil {
		return false
	}

	ty := *attrSch.Type

	switch ty {
	case cty.Number:
		dval = cty.Zero
	case cty.Bool:
		dval = cty.False
	case cty.String:
		dval = cty.StringVal("")
	default:
		switch {
		case ty.IsListType():
			dval = cty.ListValEmpty(ty.ElementType())
		case ty.IsSetType():
			dval = cty.SetValEmpty(ty.ElementType())
		case ty.IsMapType():
			dval = cty.MapValEmpty(ty.ElementType())
		}
	}

	if dval.IsNull() {
		return false
	}

	return aval.Equals(dval).True()
}

// attributeIsDefaultValue reports whether the given (parsed and normalized)
// attribute value equals the schema-defined default value.
// Note that null is the "default" default value when no default is defined.
func attributeIsDefaultValue(aval cty.Value, attrSch *tfpluginschema.SchemaAttribute) (bool, error) {
	if attrSch.Default == nil {
		if aval.IsNull() {
			return true, nil
		} else {
			return false, nil
		}
	}
	if attrSch.Type == nil {
		return false, nil
	}
	dval, err := gocty.ToCtyValue(attrSch.Default, *attrSch.Type)
	if err != nil {
		return false, fmt.Errorf("converting default value %v: %v", attrSch.Default, err)
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
