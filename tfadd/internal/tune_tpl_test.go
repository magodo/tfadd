package internal

import (
	"testing"

	"github.com/magodo/tfadd/schema"
	tfpluginschema "github.com/magodo/tfpluginschema/schema"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestTuneTpl(t *testing.T) {
	sch := schema.Schema{
		Block: &tfpluginschema.SchemaBlock{
			BlockTypes: []*tfpluginschema.SchemaNestedBlock{
				{
					TypeName: "req",
					Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
					Required: ToPtr(true),
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
	actual, err := TuneTpl(sch, []byte(input), &TuneOption{RemoveOC: true, RemoveOZAttribute: true})
	require.NoError(t, err)
	require.Equal(t, expect, string(actual))
}

func TestTuneForBlock(t *testing.T) {
	cases := []struct {
		name   string
		schema tfpluginschema.SchemaBlock
		input  string
		expect string
		ocKeep map[string]bool
	}{
		{
			name: "primary attributes only",
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:     "req",
						Type:     ToPtr(cty.Number),
						Required: true,
					},
					{
						Name:     "opt",
						Type:     ToPtr(cty.Number),
						Optional: true,
					},
					{
						Name:     "comp",
						Type:     ToPtr(cty.Number),
						Computed: true,
					},
					{
						Name:     "oc",
						Type:     ToPtr(cty.Number),
						Computed: true,
						Optional: true,
					},
				},
			},
			input: `resource "foo" "test" {
  req = 0
  opt = 1
  comp = 2
  oc = 3
}`,
			expect: `resource "foo" "test" {
  req = 0
  opt = 1
}`,
		},
		{
			name: "optional attributes with default value",
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:     "number",
						Type:     ToPtr(cty.Number),
						Optional: true,
					},
					{
						Name:     "bool",
						Type:     ToPtr(cty.Bool),
						Optional: true,
					},
					{
						Name:     "string",
						Type:     ToPtr(cty.String),
						Optional: true,
					},
					{
						Name:     "list",
						Type:     ToPtr(cty.List(cty.Number)),
						Optional: true,
					},
					{
						Name:     "set",
						Type:     ToPtr(cty.Set(cty.Number)),
						Optional: true,
					},
					{
						Name:     "map",
						Type:     ToPtr(cty.Map(cty.Number)),
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
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:     "number",
						Type:     ToPtr(cty.Number),
						Optional: true,
						Default:  1,
					},
					{
						Name:     "bool",
						Type:     ToPtr(cty.Bool),
						Optional: true,
						Default:  true,
					},
					{
						Name:     "string",
						Type:     ToPtr(cty.String),
						Optional: true,
						Default:  "default",
					},
					{
						Name:     "list",
						Type:     ToPtr(cty.List(cty.Number)),
						Optional: true,
						Default:  []interface{}{1}, // []interface{} works
					},
					{
						Name:     "set",
						Type:     ToPtr(cty.Set(cty.Number)),
						Optional: true,
						Default:  []int{1}, // []int also works
					},
					{
						Name:     "map",
						Type:     ToPtr(cty.Map(cty.Number)),
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
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:     "number",
						Type:     ToPtr(cty.Number),
						Optional: true,
					},
					{
						Name:     "bool",
						Type:     ToPtr(cty.Bool),
						Optional: true,
					},
					{
						Name:     "string",
						Type:     ToPtr(cty.String),
						Optional: true,
					},
					{
						Name:     "list",
						Type:     ToPtr(cty.List(cty.Number)),
						Optional: true,
					},
					{
						Name:     "set",
						Type:     ToPtr(cty.Set(cty.Number)),
						Optional: true,
					},
					{
						Name:     "map",
						Type:     ToPtr(cty.Map(cty.Number)),
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
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:         "attr1",
						Type:         ToPtr(cty.Number),
						Optional:     true,
						Computed:     true,
						ExactlyOneOf: []string{"attr1", "attr2"},
					},
					{
						Name:         "attr2",
						Type:         ToPtr(cty.Number),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName: "blk",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Required: ToPtr(true),
						Block: &tfpluginschema.SchemaBlock{
							Attributes: []*tfpluginschema.SchemaAttribute{
								{
									Name:         "attr1",
									Type:         ToPtr(cty.Number),
									Optional:     true,
									Computed:     true,
									ExactlyOneOf: []string{"blk.0.attr1", "blk.0.attr2"},
								},
								{
									Name:         "attr2",
									Type:         ToPtr(cty.Number),
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
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:         "attr1",
						Type:         ToPtr(cty.Number),
						Optional:     true,
						Computed:     true,
						AtLeastOneOf: []string{"attr1", "attr2"},
					},
					{
						Name:         "attr2",
						Type:         ToPtr(cty.Number),
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
			schema: tfpluginschema.SchemaBlock{
				Attributes: []*tfpluginschema.SchemaAttribute{
					{
						Name:     "attr1",
						Type:     ToPtr(cty.Number),
						Optional: true,
						Computed: true,
					},
					{
						Name:     "attr2",
						Type:     ToPtr(cty.Number),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName: "req",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Required: ToPtr(true),
					},
					{
						TypeName: "opt",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional: ToPtr(true),
					},
					{
						TypeName: "comp",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Computed: ToPtr(true),
					},
					{
						TypeName: "oc",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional: ToPtr(true),
						Computed: ToPtr(true),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName: "req",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Required: ToPtr(true),
					},
					{
						TypeName: "opt",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional: ToPtr(true),
					},
					{
						TypeName: "comp",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Computed: ToPtr(true),
					},
					{
						TypeName: "oc",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional: ToPtr(true),
						Computed: ToPtr(true),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName:     "blk1",
						Nesting:      tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional:     ToPtr(true),
						Computed:     ToPtr(true),
						ExactlyOneOf: []string{"blk1", "blk2"},
					},
					{
						TypeName:     "blk2",
						Nesting:      tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional:     ToPtr(true),
						Computed:     ToPtr(true),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName: "blk",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Required: ToPtr(true),
						Block: &tfpluginschema.SchemaBlock{
							BlockTypes: []*tfpluginschema.SchemaNestedBlock{
								{
									TypeName:     "blk1",
									Nesting:      tfpluginschema.SchemaNestedBlockNestingModeSingle,
									Optional:     ToPtr(true),
									Computed:     ToPtr(true),
									ExactlyOneOf: []string{"blk.0.blk1", "blk.0.blk2"},
								},
								{
									TypeName:     "blk2",
									Nesting:      tfpluginschema.SchemaNestedBlockNestingModeSingle,
									Optional:     ToPtr(true),
									Computed:     ToPtr(true),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName:     "blk1",
						Nesting:      tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional:     ToPtr(true),
						Computed:     ToPtr(true),
						AtLeastOneOf: []string{"blk1", "blk2"},
					},
					{
						TypeName:     "blk2",
						Nesting:      tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional:     ToPtr(true),
						Computed:     ToPtr(true),
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
			schema: tfpluginschema.SchemaBlock{
				BlockTypes: []*tfpluginschema.SchemaNestedBlock{
					{
						TypeName: "blk1",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional: ToPtr(true),
						Computed: ToPtr(true),
					},
					{
						TypeName: "blk2",
						Nesting:  tfpluginschema.SchemaNestedBlockNestingModeSingle,
						Optional: ToPtr(true),
						Computed: ToPtr(true),
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
			actual, err := TuneTpl(schema.Schema{Block: &c.schema}, []byte(c.input), &TuneOption{RemoveOC: true, RemoveOZAttribute: true, OCToKeep: c.ocKeep})
			require.NoError(t, err)
			require.Equal(t, c.expect, string(actual))
		})
	}
}

func ToPtr[T any](v T) *T {
	return &v
}
