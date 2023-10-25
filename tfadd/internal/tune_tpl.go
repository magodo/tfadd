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

func TuneTpl(sch schema.Schema, tpl []byte, rt string, ocToKeep map[string]bool) ([]byte, error) {
	f, diag := hclwrite.ParseConfig(tpl, "", hcl.InitialPos)
	if diag.HasErrors() {
		return nil, fmt.Errorf("parsing the generated template for %s: %s", rt, diag.Error())
	}
	rb := f.Body().Blocks()[0].Body()

	rb.RemoveAttribute("id")
	rb.RemoveBlock(rb.FirstMatchingBlock("timeouts", nil))

	if err := tuneForBlock(rb, sch.Block, nil, ocToKeep); err != nil {
		return nil, err
	}
	return f.Bytes(), nil
}

func tuneForBlock(rb *hclwrite.Body, sch *tfpluginschema.Block, parentAttrNames []string, ocToKeep map[string]bool) error {
	for attrName, attrVal := range rb.Attributes() {
		schAttr, ok := sch.Attributes[attrName]
		if !ok {
			// This might because the provider under used is a newer one than the version where we ingest the schema information.
			// This might happen when the user has a newer version provider installed in its local fs, and has set the "dev_overrides" for that provider.
			// We simply remove that attribute from the config.
			rb.RemoveAttribute(attrName)
			continue
		}
		if schAttr.Required {
			continue
		}

		if schAttr.Computed {
			if schAttr.Optional {
				if len(schAttr.ExactlyOneOf) != 0 {
					// For O+C attribute that has "ExactlyOneOf" constraint, keeps the first one in alphabetic order.
					l := make([]string, len(schAttr.ExactlyOneOf))
					copy(l, schAttr.ExactlyOneOf)
					sort.Strings(l)

					addrs := append(parentAttrNames, attrName)
					if l[0] != strings.Join(addrs, ".0.") {
						rb.RemoveAttribute(attrName)
						continue
					}
				} else if len(schAttr.AtLeastOneOf) == 0 {
					// For O+C attribute that has "AtLeastOneOf" constraint, or is explicitly specified, keep it.
					if !(len(ocToKeep) != 0 && ocToKeep[attrName]) {
						rb.RemoveAttribute(attrName)
						continue
					}
				}
			} else {
				rb.RemoveAttribute(attrName)
				continue
			}
		}

		// For optional only attributes, remove it from the output config if it either holds the default value or is null.
		aval, err := attrValue(attrName, attrVal)
		if err != nil {
			return err
		}
		if aval.IsNull() {
			rb.RemoveAttribute(attrName)
			continue
		}

		// Non null attribute, continue checking whether it equals to the default value.
		var dval cty.Value
		switch schAttr.Type {
		case cty.Number:
			dval = cty.Zero
		case cty.Bool:
			dval = cty.False
		case cty.String:
			dval = cty.StringVal("")
		default:
			if schAttr.Type.IsListType() {
				dval = cty.ListValEmpty(schAttr.Type.ElementType())
				if len(aval.AsValueSlice()) == 0 {
					aval = dval
				} else {
					aval = cty.ListVal(aval.AsValueSlice())
				}
				break
			}
			if schAttr.Type.IsSetType() {
				dval = cty.SetValEmpty(schAttr.Type.ElementType())
				if len(aval.AsValueSlice()) == 0 {
					aval = dval
				} else {
					aval = cty.SetVal(aval.AsValueSlice())
				}
				break
			}
			if schAttr.Type.IsMapType() {
				dval = cty.MapValEmpty(schAttr.Type.ElementType())
				if len(aval.AsValueMap()) == 0 {
					aval = dval
				} else {
					aval = cty.MapVal(aval.AsValueMap())
				}
				break
			}
		}
		if schAttr.Default != nil {
			var err error
			dval, err = gocty.ToCtyValue(schAttr.Default, schAttr.Type)
			if err != nil {
				return fmt.Errorf("converting cty value %v to Go: %v", schAttr.Default, err)
			}
		}
		if aval.Equals(dval).True() {
			rb.RemoveAttribute(attrName)
			continue
		}
	}

	for _, blkVal := range rb.Blocks() {
		scht := sch.NestedBlocks[blkVal.Type()]

		if scht.Computed {
			if scht.Optional {
				if len(scht.ExactlyOneOf) != 0 {
					// For O+C block that has "ExactlyOneOf" constraint, keeps the first one in alphabetic order.
					l := make([]string, len(scht.ExactlyOneOf))
					copy(l, scht.ExactlyOneOf)
					sort.Strings(l)

					addrs := append(parentAttrNames, blkVal.Type())
					if l[0] != strings.Join(addrs, ".0.") {
						rb.RemoveBlock(blkVal)
						continue
					}
				} else if len(scht.AtLeastOneOf) == 0 {
					// For O+blocks attribute that has "AtLeastOneOf" constraint, or is explicitly specified, keep it.
					if !(len(ocToKeep) != 0 && ocToKeep[blkVal.Type()]) {
						rb.RemoveBlock(blkVal)
						continue
					}
					continue
				}
			} else {
				// Computed only
				rb.RemoveBlock(blkVal)
				continue
			}
		}

		if err := tuneForBlock(blkVal.Body(), scht.Block, append(parentAttrNames, blkVal.Type()), nil); err != nil {
			return err
		}
	}
	return nil
}

func attrValue(attrName string, attr *hclwrite.Attribute) (cty.Value, error) {
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
