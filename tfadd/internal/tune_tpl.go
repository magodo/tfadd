package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/magodo/tfadd/schema/legacy"
	"github.com/magodo/tfadd/tfadd/internal/graph"

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

	if err := TrimHCL(rb, sch.Block); err != nil {
		return nil, err
	}

	if !opt.IgnoreAttrConstraints {
		// Construct the graph only for optional attributes.
		g, err := BuildGraph(rb, sch.Block)
		if err != nil {
			return nil, fmt.Errorf("building graph: %v", err)
		}

		// Remove the graph nodes that are deprecated (considering "R")

		// Enumerate solutions
	}

	return f.Bytes(), nil
}

func BuildGraph(rb *hclwrite.Body, sch *legacy.SchemaBlock) (*graph.Graph, error) {
	g := graph.Graph{
		Nodes: map[string]*graph.Node{},
	}

	if err := addNodes(rb, &g, sch, nil); err != nil {
		return nil, fmt.Errorf("adding nodes: %v", err)
	}

	// add edges
	return &g, nil
}

// Add nodes to the graph, which are Optional attributes, whose value are not null/zero/default.
func addNodes(rb *hclwrite.Body, g *graph.Graph, sch *legacy.SchemaBlock, parentAttrNames []string) error {
	for attrName, attrVal := range rb.Attributes() {
		schAttr, ok := sch.Attributes[attrName]
		if !ok {
			// This should have already been removed in a prior step.
			continue
		}

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

		// Not add the node to graph if it is null.
		if aval.IsNull() {
			continue
		}

		// Not add the node to graph if it is of zero or default value.
		isZero, err := isZeroValue(aval, *schAttr)
		if err != nil {
			return err
		}
		if isZero {
			continue
		}

		addr := append(parentAttrNames, attrName)
		g.AddNode(strings.Join(addr, "."))
	}

	blkCounter := map[string]int{}
	for _, blkVal := range rb.Blocks() {
		t := blkVal.Type()
		scht, ok := sch.NestedBlocks[t]

		if !ok {
			// This should have already been removed in a prior step.
			continue
		}

		if !scht.Optional {
			continue
		}

		// Add parent.foo to graph
		addr := append(parentAttrNames, t)
		g.AddNode(strings.Join(addr, "."))

		index := strconv.Itoa(blkCounter[t])
		blkCounter[t] = blkCounter[t] + 1

		// Add parent.foo.x (x is a number like "0") to graph
		addr = append(addr, index)
		g.AddNode(strings.Join(addr, "."))

		if err := addNodes(blkVal.Body(), g, scht.Block, addr); err != nil {
			return err
		}
	}

	return nil
}

func addEdges(g *graph.Graph, sch *legacy.SchemaBlock, parentAttrNames []string) {
	for name, value := range sch.Attributes {

	}
}

// TrimHCL removes the computed only attributes.
func TrimHCL(rb *hclwrite.Body, sch *legacy.SchemaBlock) error {
	for attrName := range rb.Attributes() {
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
	}

	for _, blkVal := range rb.Blocks() {
		scht, ok := sch.NestedBlocks[blkVal.Type()]

		if !ok {
			// This might because the provider under used is a newer one than the version where we ingest the schema information.
			// This might happen when the user has a newer version provider installed in its local fs, and has set the "dev_overrides" for that provider.
			// We simply remove that attribute from the config.
			rb.RemoveBlock(blkVal)
			continue
		}

		// Computed only
		if scht.Computed && !scht.Optional {
			rb.RemoveBlock(blkVal)
			continue
		}

		// Optional/Required
		if err := TrimHCL(blkVal.Body(), scht.Block); err != nil {
			return err
		}
	}
	return nil
}

func isZeroValue(value cty.Value, sch legacy.SchemaAttribute) (bool, error) {
	var dval cty.Value
	switch sch.AttributeType {
	case cty.Number:
		dval = cty.Zero
	case cty.Bool:
		dval = cty.False
	case cty.String:
		dval = cty.StringVal("")
	default:
		if sch.AttributeType.IsListType() {
			dval = cty.ListValEmpty(sch.AttributeType.ElementType())
			if len(value.AsValueSlice()) == 0 {
				value = dval
			} else {
				value = cty.ListVal(value.AsValueSlice())
			}
			if sch.AttributeType.IsSetType() {
				dval = cty.SetValEmpty(sch.AttributeType.ElementType())
				if len(value.AsValueSlice()) == 0 {
					value = dval
				} else {
					value = cty.SetVal(value.AsValueSlice())
				}
				break
			}
			if sch.AttributeType.IsMapType() {
				dval = cty.MapValEmpty(sch.AttributeType.ElementType())
				if len(value.AsValueMap()) == 0 {
					value = dval
				} else {
					value = cty.MapVal(value.AsValueMap())
				}
				break
			}
		}
		if sch.Default != nil {
			var err error
			dval, err = gocty.ToCtyValue(sch.Default, sch.AttributeType)
			if err != nil {
				return false, fmt.Errorf("converting cty value %v to Go: %v", sch.Default, err)
			}
		}
	}
	return value.Equals(dval).True(), nil
}
