package internal

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	addr2 "github.com/magodo/tfadd/addr"
	"github.com/zclconf/go-cty/cty/function/stdlib"

	"github.com/hashicorp/hcl/v2/hclwrite"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
	"github.com/zclconf/go-cty/cty"
)

func StateToTpl(r *tfstate.StateResource, schema *tfjson.SchemaBlock) ([]byte, error) {
	var buf strings.Builder
	addr, err := addr2.ParseResourceAddr(r.Address)
	if err != nil {
		return nil, fmt.Errorf("parsing resource address: %v", err)
	}
	buf.WriteString(fmt.Sprintf("resource %q %q {\n", addr.Type, addr.Name))

	// Special handling on attribute "id" to make it a Computed only attribute. This is mainly for the provider that is using the plugin sdk v2, where it is set to be O+C.
	schema.Attributes["id"] = &tfjson.SchemaAttribute{
		AttributeType: cty.String,
		Computed:      true,
	}

	if err := addAttributes(&buf, r.Value, schema.Attributes, 2); err != nil {
		return nil, err
	}
	if err := addBlocks(&buf, r.Value, schema.NestedBlocks, 2); err != nil {
		return nil, err
	}
	buf.WriteString("}\n")
	return hclwrite.Format([]byte(buf.String())), nil
}

func addAttributes(buf *strings.Builder, stateVal cty.Value, attrs map[string]*tfjson.SchemaAttribute, indent int) error {
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
		if attrS.AttributeNestedType != nil {
			// This shouldn't happen in real usage; state always has all values (set
			// to null as needed), but it protects against panics in tests (and any
			// really weird and unlikely cases).
			if !stateVal.Type().HasAttribute(name) {
				continue
			}
			nestedVal := stateVal.GetAttr(name)
			if err := addAttributeNestedTypeAttributes(buf, name, attrS, nestedVal, indent); err != nil {
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

			buf.WriteString(strings.Repeat(" ", indent))
			buf.WriteString(fmt.Sprintf("%s = ", name))
			tok := hclwrite.TokensForValue(val)
			// use jsonencode if val is valid json object
			bs := tok.Bytes()
			if attrS.AttributeType.Equals(cty.String) {
				if unquoted, err := strconv.Unquote(string(bs)); err == nil && len(unquoted) > 0 {
					if (unquoted[0] == '{' || unquoted[0] == '[') && json.Valid([]byte(unquoted)) {
						if decodeVal, err := stdlib.JSONDecode(val); err == nil {
							bs2 := hclwrite.TokensForValue(decodeVal).Bytes()
							bs = append([]byte("jsonencode("), append(bs2, ')')...)
						}
					}
				}
			}
			buf.Write(bs)

			buf.WriteString("\n")
		}
	}
	return nil
}

func addAttributeNestedTypeAttributes(buf *strings.Builder, name string, schema *tfjson.SchemaAttribute, stateVal cty.Value, indent int) error {
	if stateVal.IsNull() {
		return nil
	}
	buf.WriteString(strings.Repeat(" ", indent))
	buf.WriteString(fmt.Sprintf("%s = ", name))
	switch schema.AttributeNestedType.NestingMode {
	case tfjson.SchemaNestingModeSingle:
		buf.WriteString("{\n")

		if err := addAttributes(buf, stateVal, schema.AttributeNestedType.Attributes, indent+2); err != nil {
			return err
		}
		buf.WriteString("}\n")
		return nil

	case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
		buf.WriteString("[\n")

		listVals := ctyCollectionValues(stateVal)
		for i := range listVals {
			buf.WriteString(strings.Repeat(" ", indent+2))

			buf.WriteString("{\n")
			if err := addAttributes(buf, listVals[i], schema.AttributeNestedType.Attributes, indent+4); err != nil {
				return err
			}
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString("},\n")
		}
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString("]\n")
		return nil

	case tfjson.SchemaNestingModeMap:
		buf.WriteString("{\n")

		vals := stateVal.AsValueMap()
		keys := make([]string, 0, len(vals))
		for key := range vals {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(fmt.Sprintf("%s = {", key))

			buf.WriteString("\n")
			if err := addAttributes(buf, vals[key], schema.AttributeNestedType.Attributes, indent+4); err != nil {
				return err
			}
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString("}\n")
		}
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString("}\n")
		return nil

	default:
		// This should not happen, the above should be exhaustive.
		return fmt.Errorf("unsupported NestingMode %s", schema.AttributeNestedType.NestingMode)
	}
}

func addBlocks(buf *strings.Builder, stateVal cty.Value, blocks map[string]*tfjson.SchemaBlockType, indent int) error {
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
		if err := addNestedBlock(buf, name, blockS, blockVal, indent); err != nil {
			return err
		}
	}

	return nil
}

func addNestedBlock(buf *strings.Builder, name string, schema *tfjson.SchemaBlockType, stateVal cty.Value, indent int) error {
	if stateVal.IsNull() {
		return nil
	}
	switch schema.NestingMode {
	case tfjson.SchemaNestingModeSingle, tfjson.SchemaNestingModeGroup:
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString(fmt.Sprintf("%s {", name))

		buf.WriteString("\n")
		if err := addAttributes(buf, stateVal, schema.Block.Attributes, indent+2); err != nil {
			return err
		}
		if err := addBlocks(buf, stateVal, schema.Block.NestedBlocks, indent+2); err != nil {
			return err
		}
		buf.WriteString("}\n")
		return nil
	case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
		listVals := ctyCollectionValues(stateVal)
		for i := range listVals {
			buf.WriteString(strings.Repeat(" ", indent))
			buf.WriteString(fmt.Sprintf("%s {\n", name))
			if err := addAttributes(buf, listVals[i], schema.Block.Attributes, indent+2); err != nil {
				return err
			}
			if err := addBlocks(buf, listVals[i], schema.Block.NestedBlocks, indent+2); err != nil {
				return err
			}
			buf.WriteString("}\n")
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
			buf.WriteString(strings.Repeat(" ", indent))
			buf.WriteString(fmt.Sprintf("%s %q {", name, key))
			buf.WriteString("\n")

			if err := addAttributes(buf, vals[key], schema.Block.Attributes, indent+2); err != nil {
				return err
			}
			if err := addBlocks(buf, vals[key], schema.Block.NestedBlocks, indent+2); err != nil {
				return err
			}
			buf.WriteString(strings.Repeat(" ", indent))
			buf.WriteString("}\n")
		}
		return nil
	default:
		// This should not happen, the above should be exhaustive.
		return fmt.Errorf("unsupported NestingMode %s", schema.NestingMode)
	}
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
