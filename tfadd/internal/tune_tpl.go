package internal

import (
	"fmt"
	"github.com/magodo/tfadd/schema/legacy"
	"sort"
	"strings"

	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func TuneTpl(tpl []byte) ([]byte, error) {
	f, diag := hclwrite.ParseConfig(tpl, "", hcl.InitialPos)
	if diag.HasErrors() {
		return nil, fmt.Errorf("parsing the generated template: %s", diag.Error())
	}
	rb := f.Body().Blocks()[0].Body()

	rb.RemoveAttribute("id")
	rb.RemoveBlock(rb.FirstMatchingBlock("timeouts", nil))

	return f.Bytes(), nil
}

func TuneTplWithSchema(tpl []byte, sch legacy.Schema, opt *TuneOption) ([]byte, error) {
	f, diag := hclwrite.ParseConfig(tpl, "", hcl.InitialPos)
	if diag.HasErrors() {
		return nil, fmt.Errorf("parsing the generated template: %s", diag.Error())
	}
	rb := f.Body().Blocks()[0].Body()

	rb.RemoveAttribute("id")
	rb.RemoveBlock(rb.FirstMatchingBlock("timeouts", nil))

	if err := tuneForBlock(rb, sch.Block, nil, opt); err != nil {
		return nil, err
	}
	return f.Bytes(), nil
}

func tuneForBlock(rb *hclwrite.Body, sch *legacy.SchemaBlock, parentAttrNames []string, opt *TuneOption) error {
	if opt == nil {
		opt = &TuneOption{}
	}
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

		// Computed only
		if schAttr.Computed && !schAttr.Optional {
			rb.RemoveAttribute(attrName)
			continue
		}

		// Optional
		if !opt.IgnoreAttrConstraints {
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
			} else if schAttr.Computed && len(schAttr.AtLeastOneOf) == 0 {
				// For O+C attribute that has "AtLeastOneOf" constraint, keep it.
				rb.RemoveAttribute(attrName)
				continue
			}
		}

		// For optional attributes, remove it from the output config if it is null.
		attrExpr, diags := hclwrite.ParseConfig(attrVal.BuildTokens(nil).Bytes(), "generate_attr", hcl.InitialPos)
		if diags.HasErrors() {
			return fmt.Errorf(`building attribute %q attribute: %s`, attrName, diags.Error())
		}
		attrValLit := attrExpr.Body().GetAttribute(attrName).Expr().BuildTokens(nil).Bytes()
		dexpr, diags := hclsyntax.ParseExpression(attrValLit, "", hcl.InitialPos)
		if diags.HasErrors() {
			return fmt.Errorf(`parsing HCL expression %q: %s`, string(attrValLit), diags.Error())
		}
		aval, diags := dexpr.Value(nil)
		if diags.HasErrors() {
			return fmt.Errorf(`evaluating value of HCL expression %q: %s`, string(attrValLit), diags.Error())
		}

		if aval.IsNull() {
			rb.RemoveAttribute(attrName)
			continue
		}

		if !opt.KeepDefaultValueAttrs {
			// Non null attribute, continue checking whether it equals to the default value.
			var dval cty.Value
			switch schAttr.AttributeType {
			case cty.Number:
				dval = cty.Zero
			case cty.Bool:
				dval = cty.False
			case cty.String:
				dval = cty.StringVal("")
			default:
				if schAttr.AttributeType.IsListType() {
					dval = cty.ListValEmpty(schAttr.AttributeType.ElementType())
					if len(aval.AsValueSlice()) == 0 {
						aval = dval
					} else {
						aval = cty.ListVal(aval.AsValueSlice())
					}
					break
				}
				if schAttr.AttributeType.IsSetType() {
					dval = cty.SetValEmpty(schAttr.AttributeType.ElementType())
					if len(aval.AsValueSlice()) == 0 {
						aval = dval
					} else {
						aval = cty.SetVal(aval.AsValueSlice())
					}
					break
				}
				if schAttr.AttributeType.IsMapType() {
					dval = cty.MapValEmpty(schAttr.AttributeType.ElementType())
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
				dval, err = gocty.ToCtyValue(schAttr.Default, schAttr.AttributeType)
				if err != nil {
					return fmt.Errorf("converting cty value %v to Go: %v", schAttr.Default, err)
				}
			}
			if aval.Equals(dval).True() {
				rb.RemoveAttribute(attrName)
				continue
			}
		}
	}

	for _, blkVal := range rb.Blocks() {
		scht := sch.NestedBlocks[blkVal.Type()]

		// Computed only
		if scht.Computed && !scht.Optional {
			rb.RemoveBlock(blkVal)
			continue
		}

		// Optional
		if !opt.IgnoreAttrConstraints {
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
			} else if scht.Computed && len(scht.AtLeastOneOf) == 0 {
				// For O+C block that has "AtLeastOneOf" constraint, keep it.
				rb.RemoveBlock(blkVal)
				continue
			}
		}

		if err := tuneForBlock(blkVal.Body(), scht.Block, append(parentAttrNames, blkVal.Type()), opt); err != nil {
			return err
		}
	}
	return nil
}
