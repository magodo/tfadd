package internal

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
	"github.com/zclconf/go-cty/cty"
)

func Test_StateToTpl(t *testing.T) {
	res := &tfstate.StateResource{
		Address: "test_instance.foo",
		Value: cty.ObjectVal(map[string]cty.Value{
			"id":  cty.StringVal("some-id"),
			"ami": cty.StringVal("ami-123456789"),
			"disks": cty.ObjectVal(map[string]cty.Value{
				"mount_point": cty.StringVal("/mnt/foo"),
				"size":        cty.StringVal("50GB"),
			}),
			"list_str": cty.StringVal(`[1, 2, 3]`),
			"foo_list": cty.ListVal([]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(2),
				cty.NumberIntVal(3),
			}),
			"foo_json":     cty.StringVal(`{"foo": "bar", "block1": { "foo2": "bar2"}}`),
			"s_number":     cty.NumberIntVal(123),
			"s_bool":       cty.BoolVal(true),
			"s_string":     cty.StringVal("abc"),
			"s_list":       cty.ListVal([]cty.Value{cty.StringVal("abc")}),
			"s_set":        cty.SetVal([]cty.Value{cty.StringVal("abc")}),
			"s_map":        cty.MapVal(map[string]cty.Value{"foo": cty.StringVal("abc")}),
			"s_blk_single": cty.ObjectVal(map[string]cty.Value{"foo": cty.StringVal("abc")}),
			"s_blk_map":    cty.MapVal(map[string]cty.Value{"bar": cty.ObjectVal(map[string]cty.Value{"foo": cty.StringVal("abc")})}),
			"s_blk_list":   cty.ListVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{"foo": cty.StringVal("abc")})}),
			"s_blk_set":    cty.SetVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{"foo": cty.StringVal("abc")})}),
		}),
	}

	for _, tt := range []struct {
		name   string
		opt    *Option
		expect string
	}{
		{
			name: "Default behavior",
			expect: `resource "test_instance" "foo" {
  ami = "ami-123456789"
  disks = {
    mount_point = "/mnt/foo"
    size        = "50GB"
  }
  foo_json = jsonencode({
    block1 = {
      foo2 = "bar2"
    }
    foo = "bar"
  })
  foo_list = [1, 2, 3]
  list_str = jsonencode([1, 2, 3])
  s_blk_list = [
    {
      foo = "abc"
    },
  ]
  s_blk_map = {
    bar = {
      foo = "abc"
    }
  }
  s_blk_set = [
    {
      foo = "abc"
    },
  ]
  s_blk_single = {
    foo = "abc"
  }
  s_bool = true
  s_list = ["abc"]
  s_map = {
    foo = "abc"
  }
  s_number = 123
  s_set    = ["abc"]
  s_string = "abc"
}
`,
		},
		{
			name: "SensitveMasked",
			opt:  &Option{MaskSensitive: true},
			expect: `resource "test_instance" "foo" {
  ami = "ami-123456789"
  disks = {
    mount_point = "/mnt/foo"
    size        = "50GB"
  }
  foo_json = jsonencode({
    block1 = {
      foo2 = "bar2"
    }
    foo = "bar"
  })
  foo_list     = [1, 2, 3]
  list_str     = jsonencode([1, 2, 3])
  s_blk_list   = []    # Masked sensitive attribute
  s_blk_map    = {}    # Masked sensitive attribute
  s_blk_set    = []    # Masked sensitive attribute
  s_blk_single = {}    # Masked sensitive attribute
  s_bool       = false # Masked sensitive attribute
  s_list       = []    # Masked sensitive attribute
  s_map        = {}    # Masked sensitive attribute
  s_number     = 0     # Masked sensitive attribute
  s_set        = []    # Masked sensitive attribute
  s_string     = ""    # Masked sensitive attribute
}
`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			b, err := StateToTpl(res, addTestSchema(tfjson.SchemaNestingModeSingle), tt.opt)
			if err != nil {
				t.Fatal(err.Error())
			}
			if string(b) != tt.expect {
				t.Errorf("wrong result: %s", cmp.Diff(tt.expect, string(b)))
			}
		})
	}
}

func TestAdd_addAttributes(t *testing.T) {
	attrs := map[string]*tfjson.SchemaAttribute{
		"ami": {
			AttributeType: cty.Number,
			Optional:      true,
		},
		"boot_disk": {
			AttributeType: cty.String,
			Optional:      true,
		},
		"password": {
			AttributeType: cty.String,
			Optional:      true,
			Sensitive:     true,
		},
		"tags": {
			AttributeType: cty.Map(cty.String),
			Optional:      true,
		},
		"locations": {
			AttributeType: cty.List(cty.String),
			Optional:      true,
		},
		"ids": {
			AttributeType: cty.Set(cty.String),
			Optional:      true,
		},
		"disks": {
			AttributeNestedType: &tfjson.SchemaNestedAttributeType{
				NestingMode: tfjson.SchemaNestingModeSingle,
				Attributes: map[string]*tfjson.SchemaAttribute{
					"size": {
						AttributeType: cty.Number,
						Optional:      true,
					},
					"mount_point": {
						AttributeType: cty.String,
						Optional:      true,
					},
				},
			},
			Optional: true,
		},
	}

	tests := map[string]struct {
		attrs    map[string]*tfjson.SchemaAttribute
		val      cty.Value
		expected string
	}{
		"empty returns nil": {
			map[string]*tfjson.SchemaAttribute{},
			cty.NilVal,
			"",
		},
		"mixed attributes": {
			attrs,
			cty.ObjectVal(map[string]cty.Value{
				"ami":       cty.NumberIntVal(123456),
				"boot_disk": cty.NullVal(cty.String),
				"password":  cty.StringVal("i am secret"),
				"tags": cty.MapVal(map[string]cty.Value{
					"foo": cty.StringVal("bar"),
				}),
				"ids":       cty.SetVal([]cty.Value{cty.StringVal("999")}),
				"locations": cty.ListVal([]cty.Value{cty.StringVal("Shanghai")}),
				"disks": cty.ObjectVal(map[string]cty.Value{
					"size":        cty.NumberIntVal(50),
					"mount_point": cty.NullVal(cty.String),
				}),
			}),
			`ami = 123456
disks = {
  size = 50
}
ids = ["999"]
locations = ["Shanghai"]
password = "i am secret"
tags = {
  foo = "bar"
}
`,
		},
		"null attributes": {
			attrs,
			cty.ObjectVal(map[string]cty.Value{
				"ami":       cty.NullVal(cty.Number),
				"boot_disk": cty.NullVal(cty.String),
				"tags":      cty.NullVal(cty.Map(cty.String)),
				"ids":       cty.NullVal(cty.Set(cty.String)),
				"locations": cty.NullVal(cty.List(cty.String)),
				"disks": cty.NullVal(cty.Object(map[string]cty.Type{
					"size":        cty.Number,
					"mount_point": cty.String,
				})),
			}),
			``,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := newConverter(nil, nil)
			if err := c.AddAttributes(test.val, test.attrs, 0); err != nil {
				t.Errorf("unexpected error")
			}
			if c.String() != test.expected {
				t.Errorf("wrong result: %s", cmp.Diff(test.expected, c.String()))
			}
		})
	}
}

func TestAdd_addBlocks(t *testing.T) {
	t.Run("NestingSingle", func(t *testing.T) {
		val := cty.ObjectVal(map[string]cty.Value{
			"root_block_device": cty.ObjectVal(map[string]cty.Value{
				"volume_type": cty.StringVal("foo"),
			}),
			"network_rules": cty.NullVal(cty.Object(map[string]cty.Type{
				"ip_address": cty.String,
			})),
		})
		schema := addTestSchema(tfjson.SchemaNestingModeSingle)
		c := newConverter(nil, nil)
		c.AddBlocks(val, schema.NestedBlocks, 0)

		expected := `root_block_device {
  volume_type = "foo"
}
`

		if !cmp.Equal(c.String(), expected) {
			t.Errorf("wrong output:\n%s", cmp.Diff(expected, c.String()))
		}
	})

	t.Run("NestingList", func(t *testing.T) {
		val := cty.ObjectVal(map[string]cty.Value{
			"root_block_device": cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"volume_type": cty.StringVal("foo"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"volume_type": cty.StringVal("bar"),
				}),
			}),
			"network_rules": cty.NullVal(
				cty.List(cty.Object(map[string]cty.Type{
					"ip_address": cty.String,
				})),
			),
		})
		schema := addTestSchema(tfjson.SchemaNestingModeList)
		c := newConverter(nil, nil)
		c.AddBlocks(val, schema.NestedBlocks, 0)

		expected := `root_block_device {
  volume_type = "foo"
}
root_block_device {
  volume_type = "bar"
}
`

		if !cmp.Equal(c.String(), expected) {
			t.Fatalf("wrong output:\n%s", cmp.Diff(expected, c.String()))
		}
	})

	t.Run("NestingSet", func(t *testing.T) {
		val := cty.ObjectVal(map[string]cty.Value{
			"root_block_device": cty.SetVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"volume_type": cty.StringVal("foo"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"volume_type": cty.StringVal("bar"),
				}),
			}),
			"network_rules": cty.NullVal(
				cty.Set(cty.Object(map[string]cty.Type{
					"ip_address": cty.String,
				})),
			),
		})
		schema := addTestSchema(tfjson.SchemaNestingModeSet)
		c := newConverter(nil, nil)
		c.AddBlocks(val, schema.NestedBlocks, 0)

		expected := `root_block_device {
  volume_type = "bar"
}
root_block_device {
  volume_type = "foo"
}
`

		if !cmp.Equal(c.String(), expected) {
			t.Fatalf("wrong output:\n%s", cmp.Diff(expected, c.String()))
		}
	})

	t.Run("NestingMap", func(t *testing.T) {
		val := cty.ObjectVal(map[string]cty.Value{
			"root_block_device": cty.MapVal(map[string]cty.Value{
				"1": cty.ObjectVal(map[string]cty.Value{
					"volume_type": cty.StringVal("foo"),
				}),
				"2": cty.ObjectVal(map[string]cty.Value{
					"volume_type": cty.StringVal("bar"),
				}),
			}),
			"network_rules": cty.NullVal(
				cty.Map(cty.Object(map[string]cty.Type{
					"ip_address": cty.String,
				})),
			),
		})
		schema := addTestSchema(tfjson.SchemaNestingModeMap)
		c := newConverter(nil, nil)
		c.AddBlocks(val, schema.NestedBlocks, 0)

		expected := `root_block_device "1" {
  volume_type = "foo"
}
root_block_device "2" {
  volume_type = "bar"
}
`

		if !cmp.Equal(c.String(), expected) {
			t.Fatalf("wrong output:\n%s", cmp.Diff(expected, c.String()))
		}
	})
}

func TestAdd_addDependency(t *testing.T) {
	t.Run("Dependency", func(t *testing.T) {
		c := newConverter(nil, nil)
		c.AddDependency([]string{"foo", "bar"}, 0)
		expected := `depends_on = [
  foo,
  bar,
]
`

		if !cmp.Equal(c.String(), expected) {
			t.Errorf("wrong output:\n%s", cmp.Diff(expected, c.String()))
		}
	})
}

func addTestSchema(nesting tfjson.SchemaNestingMode) *tfjson.SchemaBlock {
	return &tfjson.SchemaBlock{
		Attributes: map[string]*tfjson.SchemaAttribute{
			"id": {AttributeType: cty.String, Optional: true, Computed: true},
			// Attributes which are neither optional nor required should not print.
			"uuid": {AttributeType: cty.String, Computed: true},
			"ami":  {AttributeType: cty.String, Optional: true},
			"disks": {
				AttributeNestedType: &tfjson.SchemaNestedAttributeType{
					Attributes: map[string]*tfjson.SchemaAttribute{
						"mount_point": {AttributeType: cty.String, Optional: true},
						"size":        {AttributeType: cty.String, Optional: true},
					},
					NestingMode: nesting,
				},
			},
			"list_str": {AttributeType: cty.String, Optional: true},
			"foo_list": {AttributeType: cty.List(cty.Number), Optional: true},
			"foo_json": {AttributeType: cty.String, Optional: true},

			// Sensitive attributes
			"s_number": {AttributeType: cty.Number, Optional: true, Sensitive: true},
			"s_bool":   {AttributeType: cty.Bool, Optional: true, Sensitive: true},
			"s_string": {AttributeType: cty.String, Optional: true, Sensitive: true},
			"s_list":   {AttributeType: cty.List(cty.String), Optional: true, Sensitive: true},
			"s_set":    {AttributeType: cty.Set(cty.String), Optional: true, Sensitive: true},
			"s_map":    {AttributeType: cty.Map(cty.String), Optional: true, Sensitive: true},
			"s_blk_single": {
				AttributeNestedType: &tfjson.SchemaNestedAttributeType{
					NestingMode: tfjson.SchemaNestingModeSingle,
					Attributes:  map[string]*tfjson.SchemaAttribute{"foo": {AttributeType: cty.String, Optional: true}}},
				Optional:  true,
				Sensitive: true,
			},
			"s_blk_map": {
				AttributeNestedType: &tfjson.SchemaNestedAttributeType{
					NestingMode: tfjson.SchemaNestingModeMap,
					Attributes:  map[string]*tfjson.SchemaAttribute{"foo": {AttributeType: cty.String, Optional: true}}},
				Optional:  true,
				Sensitive: true,
			},
			"s_blk_list": {
				AttributeNestedType: &tfjson.SchemaNestedAttributeType{
					NestingMode: tfjson.SchemaNestingModeList,
					Attributes:  map[string]*tfjson.SchemaAttribute{"foo": {AttributeType: cty.String, Optional: true}}},
				Optional:  true,
				Sensitive: true,
			},
			"s_blk_set": {
				AttributeNestedType: &tfjson.SchemaNestedAttributeType{
					NestingMode: tfjson.SchemaNestingModeSet,
					Attributes:  map[string]*tfjson.SchemaAttribute{"foo": {AttributeType: cty.String, Optional: true}}},
				Optional:  true,
				Sensitive: true,
			},
		},
		NestedBlocks: map[string]*tfjson.SchemaBlockType{
			"root_block_device": {
				Block: &tfjson.SchemaBlock{
					Attributes: map[string]*tfjson.SchemaAttribute{
						"volume_type": {
							AttributeType: cty.String,
							Optional:      true,
							Computed:      true,
						},
					},
				},
				NestingMode: nesting,
			},
			"network_rules": {
				Block: &tfjson.SchemaBlock{
					Attributes: map[string]*tfjson.SchemaAttribute{
						"ip_address": {
							AttributeType: cty.String,
							Optional:      true,
							Computed:      true,
						},
					},
				},
				NestingMode: nesting,
				MinItems:    1,
			},
		},
	}
}
