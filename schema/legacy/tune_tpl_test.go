package legacy

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestTuneTpl(t *testing.T) {
	providerSchema := ProviderSchema{
		ResourceSchemas: map[string]*Schema{
			"foo": {
				Block: &SchemaBlock{
					NestedBlocks: map[string]*SchemaBlockType{
						"req": {
							NestingMode: NestingSingle,
							Required:    true,
						},
					},
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
	actual, err := providerSchema.TuneTpl([]byte(input), "foo")
	require.NoError(t, err)
	require.Equal(t, expect, string(actual))
}

func TestTuneForBlock(t *testing.T) {
	cases := []struct {
		name   string
		schema SchemaBlock
		input  string
		expect string
	}{
		{
			name: "primary attributes only",
			schema: SchemaBlock{
				Attributes: map[string]*SchemaAttribute{
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
			schema: SchemaBlock{
				Attributes: map[string]*SchemaAttribute{
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
			schema: SchemaBlock{
				Attributes: map[string]*SchemaAttribute{
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
			schema: SchemaBlock{
				Attributes: map[string]*SchemaAttribute{
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
			schema: SchemaBlock{
				Attributes: map[string]*SchemaAttribute{
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
			schema: SchemaBlock{
				NestedBlocks: map[string]*SchemaBlockType{
					"blk": {
						NestingMode: NestingSingle,
						Required:    true,
						Block: &SchemaBlock{
							Attributes: map[string]*SchemaAttribute{
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
			name: "Blocks",
			schema: SchemaBlock{
				NestedBlocks: map[string]*SchemaBlockType{
					"req": {
						NestingMode: NestingSingle,
						Required:    true,
					},
					"opt": {
						NestingMode: NestingSingle,
						Optional:    true,
					},
					"comp": {
						NestingMode: NestingSingle,
						Computed:    true,
					},
					"oc": {
						NestingMode: NestingSingle,
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
			schema: SchemaBlock{
				NestedBlocks: map[string]*SchemaBlockType{
					"req": {
						NestingMode: NestingSingle,
						Required:    true,
					},
					"opt": {
						NestingMode: NestingSingle,
						Optional:    true,
					},
					"comp": {
						NestingMode: NestingSingle,
						Computed:    true,
					},
					"oc": {
						NestingMode: NestingSingle,
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
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f, diag := hclwrite.ParseConfig([]byte(c.input), "", hcl.InitialPos)
			require.False(t, diag.HasErrors(), diag.Error())
			rb := f.Body().Blocks()[0].Body()
			require.NoError(t, tuneForBlock(rb, &c.schema, nil))
			require.Equal(t, c.expect, string(f.Bytes()))
		})
	}
}
