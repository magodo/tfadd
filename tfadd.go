package tfadd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
	"github.com/magodo/tfstate/terraform/jsonschema"
	"github.com/zclconf/go-cty/cty"
)

type Resource struct {
	*tfstate.StateResource
}

func (v *Resource) Add(schema *tfjson.SchemaBlock) ([]byte, error) {
	var buf strings.Builder
	addr, err := parseAddress(v.Address)
	if err != nil {
		return nil, fmt.Errorf("parsing resource address: %v", err)
	}
	buf.WriteString(fmt.Sprintf("resource %q %q {\n", addr.Type, addr.Name))
	if err := v.addAttributes(&buf, v.Value, schema.Attributes, 2); err != nil {
		return nil, err
	}
	if err := v.addBlocks(&buf, v.Value, schema.NestedBlocks, 2); err != nil {
		return nil, err
	}
	buf.WriteString("}")
	return hclwrite.Format([]byte(buf.String())), nil
}

func (v *Resource) addAttributes(buf *strings.Builder, stateVal cty.Value, attrs map[string]*tfjson.SchemaAttribute, indent int) error {
	if len(attrs) == 0 {
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
			if err := v.addAttributeNestedTypeAttributes(buf, name, attrS, stateVal, indent); err != nil {
				return err
			}
			continue
		}

		// Exclude computed-only attributes
		if attrS.Required || attrS.Optional {
			buf.WriteString(strings.Repeat(" ", indent))
			buf.WriteString(fmt.Sprintf("%s = ", name))

			var val cty.Value
			if stateVal.Type().HasAttribute(name) {
				val = stateVal.GetAttr(name)
			} else {
				val = jsonschema.SchemaAttributeEmptyValue(attrS)
			}
			val, _ = val.Unmark()
			tok := hclwrite.TokensForValue(val)
			if _, err := tok.WriteTo(buf); err != nil {
				return err
			}

			buf.WriteString("\n")
		}
	}
	return nil
}

func (v *Resource) addAttributeNestedTypeAttributes(buf *strings.Builder, name string, schema *tfjson.SchemaAttribute, stateVal cty.Value, indent int) error {
	switch schema.AttributeNestedType.NestingMode {
	case tfjson.SchemaNestingModeSingle:
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString(fmt.Sprintf("%s = {\n", name))

		nestedVal := stateVal.GetAttr(name)
		if err := v.addAttributes(buf, nestedVal, schema.AttributeNestedType.Attributes, indent+2); err != nil {
			return err
		}
		buf.WriteString("}\n")
		return nil

	case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString(fmt.Sprintf("%s = [", name))

		buf.WriteString("\n")

		listVals := ctyCollectionValues(stateVal.GetAttr(name))
		for i := range listVals {
			buf.WriteString(strings.Repeat(" ", indent+2))

			buf.WriteString("{\n")
			if err := v.addAttributes(buf, listVals[i], schema.AttributeNestedType.Attributes, indent+4); err != nil {
				return err
			}
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString("},\n")
		}
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString("]\n")
		return nil

	case tfjson.SchemaNestingModeMap:
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString(fmt.Sprintf("%s = {", name))

		buf.WriteString("\n")

		vals := stateVal.GetAttr(name).AsValueMap()
		keys := make([]string, 0, len(vals))
		for key := range vals {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(fmt.Sprintf("%s = {", key))

			buf.WriteString("\n")
			if err := v.addAttributes(buf, vals[key], schema.AttributeNestedType.Attributes, indent+4); err != nil {
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

func (v *Resource) addBlocks(buf *strings.Builder, stateVal cty.Value, blocks map[string]*tfjson.SchemaBlockType, indent int) error {
	if len(blocks) == 0 {
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
		if err := v.addNestedBlock(buf, name, blockS, blockVal, indent); err != nil {
			return err
		}
	}

	return nil
}

func (v *Resource) addNestedBlock(buf *strings.Builder, name string, schema *tfjson.SchemaBlockType, stateVal cty.Value, indent int) error {
	switch schema.NestingMode {
	case tfjson.SchemaNestingModeSingle, tfjson.SchemaNestingModeGroup:
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteString(fmt.Sprintf("%s {", name))

		buf.WriteString("\n")
		if err := v.addAttributes(buf, stateVal, schema.Block.Attributes, indent+2); err != nil {
			return err
		}
		if err := v.addBlocks(buf, stateVal, schema.Block.NestedBlocks, indent+2); err != nil {
			return err
		}
		buf.WriteString("}\n")
		return nil
	case tfjson.SchemaNestingModeList, tfjson.SchemaNestingModeSet:
		listVals := ctyCollectionValues(stateVal)
		for i := range listVals {
			buf.WriteString(strings.Repeat(" ", indent))
			buf.WriteString(fmt.Sprintf("%s {\n", name))
			if err := v.addAttributes(buf, listVals[i], schema.Block.Attributes, indent+2); err != nil {
				return err
			}
			if err := v.addBlocks(buf, listVals[i], schema.Block.NestedBlocks, indent+2); err != nil {
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

			if err := v.addAttributes(buf, vals[key], schema.Block.Attributes, indent+2); err != nil {
				return err
			}
			if err := v.addBlocks(buf, vals[key], schema.Block.NestedBlocks, indent+2); err != nil {
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

type ResourceAddr struct {
	Type string
	Name string
}

func parseAddress(addr string) (*ResourceAddr, error) {
	segs := strings.Split(addr, ".")
	if len(segs) != 2 {
		return nil, fmt.Errorf("invalid resource address found: %s", addr)
	}
	return &ResourceAddr{Type: segs[0], Name: segs[1]}, nil
}
