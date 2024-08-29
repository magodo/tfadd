package internal

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	addr2 "github.com/magodo/tfadd/addr"
	"github.com/zclconf/go-cty/cty/function/stdlib"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
	"github.com/zclconf/go-cty/cty"
)

func ProviderTpl(name string, v cty.Value, schema *tfjson.SchemaBlock) ([]byte, error) {
	c := newConverter(nil, nil)
	c.WriteString(fmt.Sprintf("provider %q {\n", name))
	if err := c.AddAttributes(v, schema.Attributes, 2); err != nil {
		return nil, err
	}
	if err := c.AddBlocks(v, schema.NestedBlocks, 2); err != nil {
		return nil, err
	}
	c.WriteString("}\n")
	return hclwrite.Format([]byte(c.String())), nil
}

func StateToTpl(r *tfstate.StateResource, schema *tfjson.SchemaBlock, opt *Option) ([]byte, error) {
	c := newConverter(nil, opt)
	addr, err := addr2.ParseResourceAddr(r.Address)
	if err != nil {
		return nil, fmt.Errorf("parsing resource address: %v", err)
	}
	c.WriteString(fmt.Sprintf("resource %q %q {\n", addr.Type, addr.Name))

	// Special handling on attribute "id" to make it a Computed only attribute. This is mainly for the provider that is using the plugin sdk v2, where it is set to be O+C.
	schema.Attributes["id"].Optional = false

	if err := c.AddAttributes(r.Value, schema.Attributes, 2); err != nil {
		return nil, err
	}
	if err := c.AddBlocks(r.Value, schema.NestedBlocks, 2); err != nil {
		return nil, err
	}
	c.AddDependency(r.DependsOn, 2)
	c.WriteString("}\n")
	return hclwrite.Format([]byte(c.String())), nil
}

type converter struct {
	buf *strings.Builder
	opt Option
}

func newConverter(buf *strings.Builder, opt *Option) converter {
	if buf == nil {
		buf = &strings.Builder{}
	}
	if opt == nil {
		opt = &Option{}
	}
	return converter{
		buf: buf,
		opt: *opt,
	}
}

func (c converter) WriteString(s string) (int, error) {
	return c.buf.WriteString(s)
}

func (c converter) String() string {
	return c.buf.String()
}

func (c converter) AddAttributes(stateVal cty.Value, attrs map[string]*tfjson.SchemaAttribute, indent int) error {
	if len(attrs) == 0 || stateVal.IsNull() {
		return nil
	}

	keys := make([]string, 0, len(attrs))
	for k := range attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		name := keys[i]
		attrS := attrs[name]

		// Optionally mask sensitive attributes
		if c.opt.MaskSensitive && attrS.Sensitive {
			if err := c.AddMaskedAttr(name, attrS, indent); err != nil {
				return err
			}
			continue
		}

		if attrS.AttributeNestedType != nil {
			// This shouldn't happen in real usage; state always has all values (set
			// to null as needed), but it protects against panics in tests (and any
			// really weird and unlikely cases).
			if !stateVal.Type().HasAttribute(name) {
				continue
			}
			nestedVal := stateVal.GetAttr(name)
			if err := c.AddAttributeNestedTypeAttributes(name, attrS, nestedVal, indent); err != nil {
				return err
			}
			continue
		}

		// Exclude computed-only attributes
		if attrS.Required || attrS.Optional {
			// This shouldn't happen in real usage; state always has all values (set
			// to null as needed), but it protects against panics in tests (and any
			// really weird and unlikely cases).
			if !stateVal.Type().HasAttribute(name) {
				continue
			}

			var val cty.Value
			val = stateVal.GetAttr(name)
			val, _ = val.Unmark()
			if val.IsNull() {
				continue
			}

			c.buf.WriteString(strings.Repeat(" ", indent))
			c.buf.WriteString(fmt.Sprintf("%s = ", name))
			tok := hclwrite.TokensForValue(val)
			// use jsonencode if val is valid json object
			bs := tok.Bytes()
			if attrS.AttributeType.Equals(cty.String) {
				if unquoted, err := strconv.Unquote(string(bs)); err == nil && len(unquoted) > 0 {
					if (unquoted[0] == '{' || unquoted[0] == '[') && json.Valid([]byte(unquoted)) {
						if decodeVal, err := stdlib.JSONDecode(val); err == nil {
							bs2 := hclwrite.TokensForValue(decodeVal).Bytes()
							// Ensure the HCL representation of the JSON is still a valid HCL
							// See: https://github.com/Azure/aztfexport/issues/557 for details.
							if _, err := hclsyntax.ParseExpression(bs2, "", hcl.InitialPos); err == nil {
								bs = append([]byte("jsonencode("), append(bs2, ')')...)
							}
						}
					}
				}
			}
			c.buf.Write(bs)

			c.buf.WriteString("\n")
		}
	}
	return nil
}

func (c converter) AddDependency(deps []string, indent int) {
	if len(deps) == 0 {
		return
	}
	c.buf.WriteString(strings.Repeat(" ", indent))
	c.buf.WriteString("depends_on = [\n")
	for _, dep := range deps {
		c.buf.WriteString(strings.Repeat(" ", indent+2) + dep + ",\n")
	}
	c.buf.WriteString(strings.Repeat(" ", indent))
	c.buf.WriteString("]\n")
}

func (c converter) AddAttributeNestedTypeAttributes(name string, schema *tfjson.SchemaAttribute, stateVal cty.Value, indent int) error {
	if stateVal.IsNull() {
		return nil
	}
	c.buf.WriteString(strings.Repeat(" ", indent))
	c.buf.WriteString(fmt.Sprintf("%s = ", name))
	switch schema.AttributeNestedType.NestingMode {
	case tfjson.SchemaNestingModeSingle:
		c.buf.WriteString("{\n")

		if err := c.AddAttributes(stateVal, schema.AttributeNestedType.Attributes, indent+2); err != nil {
			return err
		}
		c.buf.WriteString("}\n")
		return nil

	case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
		c.buf.WriteString("[\n")

		listVals := ctyCollectionValues(stateVal)
		for i := range listVals {
			c.buf.WriteString(strings.Repeat(" ", indent+2))

			c.buf.WriteString("{\n")
			if err := c.AddAttributes(listVals[i], schema.AttributeNestedType.Attributes, indent+4); err != nil {
				return err
			}
			c.buf.WriteString(strings.Repeat(" ", indent+2))
			c.buf.WriteString("},\n")
		}
		c.buf.WriteString(strings.Repeat(" ", indent))
		c.buf.WriteString("]\n")
		return nil

	case tfjson.SchemaNestingModeMap:
		c.buf.WriteString("{\n")

		vals := stateVal.AsValueMap()
		keys := make([]string, 0, len(vals))
		for key := range vals {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			c.buf.WriteString(strings.Repeat(" ", indent+2))
			c.buf.WriteString(fmt.Sprintf("%s = {", key))

			c.buf.WriteString("\n")
			if err := c.AddAttributes(vals[key], schema.AttributeNestedType.Attributes, indent+4); err != nil {
				return err
			}
			c.buf.WriteString(strings.Repeat(" ", indent+2))
			c.buf.WriteString("}\n")
		}
		c.buf.WriteString(strings.Repeat(" ", indent))
		c.buf.WriteString("}\n")
		return nil

	default:
		// This should not happen, the above should be exhaustive.
		return fmt.Errorf("unsupported NestingMode %s", schema.AttributeNestedType.NestingMode)
	}
}

func (c converter) AddBlocks(stateVal cty.Value, blocks map[string]*tfjson.SchemaBlockType, indent int) error {
	if len(blocks) == 0 || stateVal.IsNull() {
		return nil
	}

	names := make([]string, 0, len(blocks))
	for k := range blocks {
		names = append(names, k)
	}
	sort.Strings(names)

	for _, name := range names {
		blockS := blocks[name]
		// This shouldn't happen in real usage; state always has all values (set
		// to null as needed), but it protects against panics in tests (and any
		// really weird and unlikely cases).
		if !stateVal.Type().HasAttribute(name) {
			continue
		}
		blockVal := stateVal.GetAttr(name)
		if err := c.AddNestedBlock(name, blockS, blockVal, indent); err != nil {
			return err
		}
	}

	return nil
}

func (c converter) AddNestedBlock(name string, schema *tfjson.SchemaBlockType, stateVal cty.Value, indent int) error {
	if stateVal.IsNull() {
		return nil
	}

	// // Converting the List and Set modes that have single-element constraint to Single mode.
	// // This is how the legacy SDK defines the Single mode.
	// if schema.MaxItems == 1 &&
	// 	slices.Index([]tfjson.SchemaNestingMode{tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet}, schema.NestingMode) != -1 {
	// 	schema.NestingMode = tfjson.SchemaNestingModeSingle
	// }

	switch schema.NestingMode {
	case tfjson.SchemaNestingModeSingle, tfjson.SchemaNestingModeGroup:
		c.buf.WriteString(strings.Repeat(" ", indent))
		c.buf.WriteString(fmt.Sprintf("%s {", name))

		c.buf.WriteString("\n")
		if err := c.AddAttributes(stateVal, schema.Block.Attributes, indent+2); err != nil {
			return err
		}
		if err := c.AddBlocks(stateVal, schema.Block.NestedBlocks, indent+2); err != nil {
			return err
		}
		c.buf.WriteString("}\n")
		return nil
	case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
		listVals := ctyCollectionValues(stateVal)
		for i := range listVals {
			c.buf.WriteString(strings.Repeat(" ", indent))
			c.buf.WriteString(fmt.Sprintf("%s {\n", name))
			if err := c.AddAttributes(listVals[i], schema.Block.Attributes, indent+2); err != nil {
				return err
			}
			if err := c.AddBlocks(listVals[i], schema.Block.NestedBlocks, indent+2); err != nil {
				return err
			}
			c.buf.WriteString("}\n")
		}
		return nil
	case tfjson.SchemaNestingModeMap:
		vals := stateVal.AsValueMap()
		keys := make([]string, 0, len(vals))
		for key := range vals {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			c.buf.WriteString(strings.Repeat(" ", indent))
			c.buf.WriteString(fmt.Sprintf("%s %q {", name, key))
			c.buf.WriteString("\n")

			if err := c.AddAttributes(vals[key], schema.Block.Attributes, indent+2); err != nil {
				return err
			}
			if err := c.AddBlocks(vals[key], schema.Block.NestedBlocks, indent+2); err != nil {
				return err
			}
			c.buf.WriteString(strings.Repeat(" ", indent))
			c.buf.WriteString("}\n")
		}
		return nil
	default:
		// This should not happen, the above should be exhaustive.
		return fmt.Errorf("unsupported NestingMode %s", schema.NestingMode)
	}
}

func (c converter) AddMaskedAttr(name string, schema *tfjson.SchemaAttribute, indent int) error {
	var v string
	if schema.AttributeNestedType != nil {
		switch schema.AttributeNestedType.NestingMode {
		case tfjson.SchemaNestingModeSingle, tfjson.SchemaNestingModeMap:
			v = "{}"
		case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
			v = "[]"
		default:
			// This should not happen, the above should be exhaustive.
			return fmt.Errorf("unsupported NestingMode %s", schema.AttributeNestedType.NestingMode)
		}
	} else {
		switch schema.AttributeType {
		case cty.Number:
			v = "0"
		case cty.Bool:
			v = "false"
		case cty.String:
			v = `""`
		default:
			switch {
			case schema.AttributeType.IsListType():
				v = "[]"
			case schema.AttributeType.IsSetType():
				v = "[]"
			case schema.AttributeType.IsMapType():
				v = "{}"
			default:
				return fmt.Errorf("unhandled attribute type: %s", schema.AttributeType.FriendlyName())
			}
		}
	}
	c.buf.WriteString(strings.Repeat(" ", indent))
	c.buf.WriteString(fmt.Sprintf("%s = %s # Masked sensitive attribute\n", name, v))
	return nil
}

func ctyCollectionValues(val cty.Value) []cty.Value {
	if !val.IsKnown() || val.IsNull() {
		return nil
	}

	var len int
	if val.IsMarked() {
		val, _ = val.Unmark()
		len = val.LengthInt()
	} else {
		len = val.LengthInt()
	}

	ret := make([]cty.Value, 0, len)
	for it := val.ElementIterator(); it.Next(); {
		_, value := it.Element()
		ret = append(ret, value)
	}

	return ret
}
