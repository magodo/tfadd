package internal

import (
	"github.com/magodo/tfadd/schema/legacy"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestTuneTpl(t *testing.T) {
	input := `resource "foo" "test" {
  id = "foo"
  timeouts {
    create = "10s"
    read = "10s"
    update = "10s"
    delete = "10s"
  }
  req {}
}`

	expect := `resource "foo" "test" {
  req {}
}`
	actual, err := TuneTpl([]byte(input))
	require.NoError(t, err)
	require.Equal(t, expect, string(actual))
}

func TestTuneForBlock(t *testing.T) {
	cases := []struct {
		name   string
		schema legacy.SchemaBlock
		input  string
		expect string
	}{
		{
			name: "primary attributes only",
			schema: legacy.SchemaBlock{
				Attributes: map[string]*legacy.SchemaAttribute{
					"req": {
						AttributeType: cty.Number,
						Required:      true,
					},
					"opt": {
						AttributeType: cty.Number,
						Optional:      true,
					},
					"comp": {
						AttributeType: cty.Number,
						Computed:      true,
					},
					"oc": {
						AttributeType: cty.Number,
						Computed:      true,
						Optional:      true,
					},
				},
			},
			input: `resource "foo" "test" {
  req = 1
  opt = 2
  comp = 3
  oc = 4
}`,
			expect: `resource "foo" "test" {
  req = 1
  opt = 2
}`,
		},
		{
			name: "optional attributes with default value",
			schema: legacy.SchemaBlock{
				Attributes: map[string]*legacy.SchemaAttribute{
					"number": {
						AttributeType: cty.Number,
						Optional:      true,
					},
					"bool": {
						AttributeType: cty.Bool,
						Optional:      true,
					},
					"string": {
						AttributeType: cty.String,
						Optional:      true,
					},
					"list": {
						AttributeType: cty.List(cty.Number),
						Optional:      true,
					},
					"set": {
						AttributeType: cty.Set(cty.Number),
						Optional:      true,
					},
					"map": {
						AttributeType: cty.Map(cty.Number),
						Optional:      true,
					},
				},
			},
			input: `resource "foo" "test" {
  number = 0
  bool = false
  string = ""
  list = []
  set = []
  map = {}
}`,
			expect: `resource "foo" "test" {
}`,
		},
		{
			name: "optional attributes with customized default value",
			schema: legacy.SchemaBlock{
				Attributes: map[string]*legacy.SchemaAttribute{
					"number": {
						AttributeType: cty.Number,
						Optional:      true,
						Default:       1,
					},
					"bool": {
						AttributeType: cty.Bool,
						Optional:      true,
						Default:       true,
					},
					"string": {
						AttributeType: cty.String,
						Optional:      true,
						Default:       "default",
					},
					"list": {
						AttributeType: cty.List(cty.Number),
						Optional:      true,
						Default:       []interface{}{1}, // []interface{} works
					},
					"set": {
						AttributeType: cty.Set(cty.Number),
						Optional:      true,
						Default:       []int{1}, // []int also works
					},
					"map": {
						AttributeType: cty.Map(cty.Number),
						Optional:      true,
						Default:       map[string]interface{}{"default": 1},
					},
				},
			},
			input: `resource "foo" "test" {
  number = 1
  bool = true
  string = "default"
  list = [1]
  set = [1]
  map = {
    default = 1
  }
}`,
			expect: `resource "foo" "test" {
}`,
		},
		{
			name: "optional attributes with null value",
			schema: legacy.SchemaBlock{
				Attributes: map[string]*legacy.SchemaAttribute{
					"number": {
						AttributeType: cty.Number,
						Optional:      true,
					},
					"bool": {
						AttributeType: cty.Bool,
						Optional:      true,
					},
					"string": {
						AttributeType: cty.String,
						Optional:      true,
					},
					"list": {
						AttributeType: cty.List(cty.Number),
						Optional:      true,
					},
					"set": {
						AttributeType: cty.Set(cty.Number),
						Optional:      true,
					},
					"map": {
						AttributeType: cty.Map(cty.Number),
						Optional:      true,
					},
				},
			},
			input: `resource "foo" "test" {
  number = null
  bool = null
  string = null
  list = null
  set = null
  map = null
}`,
			expect: `resource "foo" "test" {
}`,
		},
		{
			name: "O+C attributes that has ExactlyOneOf defined",
			schema: legacy.SchemaBlock{
				Attributes: map[string]*legacy.SchemaAttribute{
					"attr1": {
						AttributeType: cty.Number,
						Optional:      true,
						Computed:      true,
						ExactlyOneOf:  []string{"attr1", "attr2"},
					},
					"attr2": {
						AttributeType: cty.Number,
						Optional:      true,
						Computed:      true,
						ExactlyOneOf:  []string{"attr1", "attr2"},
					},
				},
			},
			input: `resource "foo" "test" {
  attr1 = 1
  attr2 = 2
}`,
			expect: `resource "foo" "test" {
  attr1 = 1
}`,
		},
		{
			name: "O+C attributes that has ExactlyOneOf defined in nested block",
			schema: legacy.SchemaBlock{
				NestedBlocks: map[string]*legacy.SchemaBlockType{
					"blk": {
						NestingMode: legacy.NestingSingle,
						Required:    true,
						Block: &legacy.SchemaBlock{
							Attributes: map[string]*legacy.SchemaAttribute{
								"attr1": {
									AttributeType: cty.Number,
									Optional:      true,
									Computed:      true,
									ExactlyOneOf:  []string{"blk.0.attr1", "blk.0.attr2"},
								},
								"attr2": {
									AttributeType: cty.Number,
									Optional:      true,
									Computed:      true,
									ExactlyOneOf:  []string{"blk.0.attr1", "blk.0.attr2"},
								},
							},
						},
					},
				},
			},
			input: `resource "foo" "test" {
  blk {
    attr1 = 1
    attr2 = 2
  }
}`,
			expect: `resource "foo" "test" {
  blk {
    attr1 = 1
  }
}`,
		},
		{
			name: "O+C attributes that has AtLeastOneOf defined",
			schema: legacy.SchemaBlock{
				Attributes: map[string]*legacy.SchemaAttribute{
					"attr1": {
						AttributeType: cty.Number,
						Optional:      true,
						Computed:      true,
						AtLeastOneOf:  []string{"attr1", "attr2"},
					},
					"attr2": {
						AttributeType: cty.Number,
						Optional:      true,
						Computed:      true,
						AtLeastOneOf:  []string{"attr1", "attr2"},
					},
				},
			},
			input: `resource "foo" "test" {
  attr1 = 1
  attr2 = 2
}`,
			expect: `resource "foo" "test" {
  attr1 = 1
  attr2 = 2
}`,
		},
		{
			name: "Blocks",
			schema: legacy.SchemaBlock{
				NestedBlocks: map[string]*legacy.SchemaBlockType{
					"req": {
						NestingMode: legacy.NestingSingle,
						Required:    true,
					},
					"opt": {
						NestingMode: legacy.NestingSingle,
						Optional:    true,
					},
					"comp": {
						NestingMode: legacy.NestingSingle,
						Computed:    true,
					},
					"oc": {
						NestingMode: legacy.NestingSingle,
						Optional:    true,
						Computed:    true,
					},
				},
			},
			input: `resource "foo" "test" {
  req {}
  opt {}
  comp {}
  oc {}
}`,
			expect: `resource "foo" "test" {
  req {}
  opt {}
}`,
		},
		{
			name: "Blocks with absent",
			schema: legacy.SchemaBlock{
				NestedBlocks: map[string]*legacy.SchemaBlockType{
					"req": {
						NestingMode: legacy.NestingSingle,
						Required:    true,
					},
					"opt": {
						NestingMode: legacy.NestingSingle,
						Optional:    true,
					},
					"comp": {
						NestingMode: legacy.NestingSingle,
						Computed:    true,
					},
					"oc": {
						NestingMode: legacy.NestingSingle,
						Optional:    true,
						Computed:    true,
					},
				},
			},
			input: `resource "foo" "test" {
  req {}
}`,
			expect: `resource "foo" "test" {
  req {}
}`,
		},
		{
			name: "O+C blocks that has ExactlyOneOf defined",
			schema: legacy.SchemaBlock{
				NestedBlocks: map[string]*legacy.SchemaBlockType{
					"blk1": {
						NestingMode:  legacy.NestingSingle,
						Optional:     true,
						Computed:     true,
						ExactlyOneOf: []string{"blk1", "blk2"},
					},
					"blk2": {
						NestingMode:  legacy.NestingSingle,
						Optional:     true,
						Computed:     true,
						ExactlyOneOf: []string{"blk1", "blk2"},
					},
				},
			},
			input: `resource "foo" "test" {
  blk1 {}
  blk2 {}
}`,
			expect: `resource "foo" "test" {
  blk1 {}
}`,
		},
		{
			name: "O+C blocks that has ExactlyOneOf defined in nested block",
			schema: legacy.SchemaBlock{
				NestedBlocks: map[string]*legacy.SchemaBlockType{
					"blk": {
						NestingMode: legacy.NestingSingle,
						Required:    true,
						Block: &legacy.SchemaBlock{
							NestedBlocks: map[string]*legacy.SchemaBlockType{
								"blk1": {
									NestingMode:  legacy.NestingSingle,
									Optional:     true,
									Computed:     true,
									ExactlyOneOf: []string{"blk.0.blk1", "blk.0.blk2"},
								},
								"blk2": {
									NestingMode:  legacy.NestingSingle,
									Optional:     true,
									Computed:     true,
									ExactlyOneOf: []string{"blk.0.blk1", "blk.0.blk2"},
								},
							},
						},
					},
				},
			},
			input: `resource "foo" "test" {
  blk {
    blk1 {}
    blk2 {}
  }
}`,
			expect: `resource "foo" "test" {
  blk {
    blk1 {}
  }
}`,
		},
		{
			name: "O+C blocks that has AtLeastOneOf defined",
			schema: legacy.SchemaBlock{
				NestedBlocks: map[string]*legacy.SchemaBlockType{
					"blk1": {
						NestingMode:  legacy.NestingSingle,
						Optional:     true,
						Computed:     true,
						AtLeastOneOf: []string{"blk1", "blk2"},
					},
					"blk2": {
						NestingMode:  legacy.NestingSingle,
						Optional:     true,
						Computed:     true,
						AtLeastOneOf: []string{"blk1", "blk2"},
					},
				},
			},
			input: `resource "foo" "test" {
  blk1 {}
  blk2 {}
}`,
			expect: `resource "foo" "test" {
  blk1 {}
  blk2 {}
}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f, diag := hclwrite.ParseConfig([]byte(c.input), "", hcl.InitialPos)
			require.False(t, diag.HasErrors(), diag.Error())
			rb := f.Body().Blocks()[0].Body()
			require.NoError(t, tuneForBlock(rb, &c.schema, nil, nil))
			require.Equal(t, c.expect, string(f.Bytes()))
		})
	}
}
