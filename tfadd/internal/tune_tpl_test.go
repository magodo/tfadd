package internal

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/magodo/tfadd/schema"
	tfpluginschema "github.com/magodo/tfpluginschema/schema"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestTuneTpl(t *testing.T) {
	sch := schema.Schema{
		Block: &tfpluginschema.Block{
			NestedBlocks: map[string]*tfpluginschema.NestedBlock{
				"req": {
					NestingMode: tfpluginschema.NestingSingle,
					Required:    true,
				},
			},
		},
	}
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
	actual, err := TuneTpl(sch, []byte(input), "foo", nil)
	require.NoError(t, err)
	require.Equal(t, expect, string(actual))
}

func TestTuneForBlock(t *testing.T) {
	cases := []struct {
		name   string
		schema tfpluginschema.Block
		input  string
		expect string
		ocKeep map[string]bool
	}{
		{
			name: "primary attributes only",
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"req": {
						Type:     cty.Number,
						Required: true,
					},
					"opt": {
						Type:     cty.Number,
						Optional: true,
					},
					"comp": {
						Type:     cty.Number,
						Computed: true,
					},
					"oc": {
						Type:     cty.Number,
						Computed: true,
						Optional: true,
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
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"number": {
						Type:     cty.Number,
						Optional: true,
					},
					"bool": {
						Type:     cty.Bool,
						Optional: true,
					},
					"string": {
						Type:     cty.String,
						Optional: true,
					},
					"list": {
						Type:     cty.List(cty.Number),
						Optional: true,
					},
					"set": {
						Type:     cty.Set(cty.Number),
						Optional: true,
					},
					"map": {
						Type:     cty.Map(cty.Number),
						Optional: true,
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
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"number": {
						Type:     cty.Number,
						Optional: true,
						Default:  1,
					},
					"bool": {
						Type:     cty.Bool,
						Optional: true,
						Default:  true,
					},
					"string": {
						Type:     cty.String,
						Optional: true,
						Default:  "default",
					},
					"list": {
						Type:     cty.List(cty.Number),
						Optional: true,
						Default:  []interface{}{1}, // []interface{} works
					},
					"set": {
						Type:     cty.Set(cty.Number),
						Optional: true,
						Default:  []int{1}, // []int also works
					},
					"map": {
						Type:     cty.Map(cty.Number),
						Optional: true,
						Default:  map[string]interface{}{"default": 1},
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
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"number": {
						Type:     cty.Number,
						Optional: true,
					},
					"bool": {
						Type:     cty.Bool,
						Optional: true,
					},
					"string": {
						Type:     cty.String,
						Optional: true,
					},
					"list": {
						Type:     cty.List(cty.Number),
						Optional: true,
					},
					"set": {
						Type:     cty.Set(cty.Number),
						Optional: true,
					},
					"map": {
						Type:     cty.Map(cty.Number),
						Optional: true,
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
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"attr1": {
						Type:         cty.Number,
						Optional:     true,
						Computed:     true,
						ExactlyOneOf: []string{"attr1", "attr2"},
					},
					"attr2": {
						Type:         cty.Number,
						Optional:     true,
						Computed:     true,
						ExactlyOneOf: []string{"attr1", "attr2"},
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
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"blk": {
						NestingMode: tfpluginschema.NestingSingle,
						Required:    true,
						Block: &tfpluginschema.Block{
							Attributes: map[string]*tfpluginschema.Attribute{
								"attr1": {
									Type:         cty.Number,
									Optional:     true,
									Computed:     true,
									ExactlyOneOf: []string{"blk.0.attr1", "blk.0.attr2"},
								},
								"attr2": {
									Type:         cty.Number,
									Optional:     true,
									Computed:     true,
									ExactlyOneOf: []string{"blk.0.attr1", "blk.0.attr2"},
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
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"attr1": {
						Type:         cty.Number,
						Optional:     true,
						Computed:     true,
						AtLeastOneOf: []string{"attr1", "attr2"},
					},
					"attr2": {
						Type:         cty.Number,
						Optional:     true,
						Computed:     true,
						AtLeastOneOf: []string{"attr1", "attr2"},
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
			name: "O+C attributes that is specified to keep",
			schema: tfpluginschema.Block{
				Attributes: map[string]*tfpluginschema.Attribute{
					"attr1": {
						Type:     cty.Number,
						Optional: true,
						Computed: true,
					},
					"attr2": {
						Type:     cty.Number,
						Optional: true,
						Computed: true,
					},
				},
			},
			ocKeep: map[string]bool{"attr1": true},
			input: `resource "foo" "test" {
  attr1 = 1
  attr2 = 2
}`,
			expect: `resource "foo" "test" {
  attr1 = 1
}`,
		},
		{
			name: "Blocks",
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"req": {
						NestingMode: tfpluginschema.NestingSingle,
						Required:    true,
					},
					"opt": {
						NestingMode: tfpluginschema.NestingSingle,
						Optional:    true,
					},
					"comp": {
						NestingMode: tfpluginschema.NestingSingle,
						Computed:    true,
					},
					"oc": {
						NestingMode: tfpluginschema.NestingSingle,
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
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"req": {
						NestingMode: tfpluginschema.NestingSingle,
						Required:    true,
					},
					"opt": {
						NestingMode: tfpluginschema.NestingSingle,
						Optional:    true,
					},
					"comp": {
						NestingMode: tfpluginschema.NestingSingle,
						Computed:    true,
					},
					"oc": {
						NestingMode: tfpluginschema.NestingSingle,
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
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"blk1": {
						NestingMode:  tfpluginschema.NestingSingle,
						Optional:     true,
						Computed:     true,
						ExactlyOneOf: []string{"blk1", "blk2"},
					},
					"blk2": {
						NestingMode:  tfpluginschema.NestingSingle,
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
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"blk": {
						NestingMode: tfpluginschema.NestingSingle,
						Required:    true,
						Block: &tfpluginschema.Block{
							NestedBlocks: map[string]*tfpluginschema.NestedBlock{
								"blk1": {
									NestingMode:  tfpluginschema.NestingSingle,
									Optional:     true,
									Computed:     true,
									ExactlyOneOf: []string{"blk.0.blk1", "blk.0.blk2"},
								},
								"blk2": {
									NestingMode:  tfpluginschema.NestingSingle,
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
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"blk1": {
						NestingMode:  tfpluginschema.NestingSingle,
						Optional:     true,
						Computed:     true,
						AtLeastOneOf: []string{"blk1", "blk2"},
					},
					"blk2": {
						NestingMode:  tfpluginschema.NestingSingle,
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
		{
			name: "O+C blocks that is specified to keep",
			schema: tfpluginschema.Block{
				NestedBlocks: map[string]*tfpluginschema.NestedBlock{
					"blk1": {
						NestingMode: tfpluginschema.NestingSingle,
						Optional:    true,
						Computed:    true,
					},
					"blk2": {
						NestingMode: tfpluginschema.NestingSingle,
						Optional:    true,
						Computed:    true,
					},
				},
			},
			ocKeep: map[string]bool{"blk1": true},
			input: `resource "foo" "test" {
  blk1 {}
  blk2 {}
}`,
			expect: `resource "foo" "test" {
  blk1 {}
}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f, diag := hclwrite.ParseConfig([]byte(c.input), "", hcl.InitialPos)
			require.False(t, diag.HasErrors(), diag.Error())
			rb := f.Body().Blocks()[0].Body()
			require.NoError(t, tuneForBlock(rb, &c.schema, nil, c.ocKeep))
			require.Equal(t, c.expect, string(f.Bytes()))
		})
	}
}
