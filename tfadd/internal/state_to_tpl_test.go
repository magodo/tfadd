package internal

import (
	"strings"
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
			"foo_list": cty.ListVal([]cty.Value{
				cty.NumberIntVal(1),
				cty.NumberIntVal(2),
				cty.NumberIntVal(3),
			}),
			"foo_json": cty.StringVal(`{"foo": "bar"}`),
		}),
	}
	b, err := StateToTpl(res, addTestSchema(tfjson.SchemaNestingModeSingle))
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := `resource "test_instance" "foo" {
  ami = "ami-123456789"
  disks = {
    mount_point = "/mnt/foo"
    size        = "50GB"
  }
  foo_json = jsonencode({
    foo = "bar"
  })
  foo_list = [1, 2, 3]
}
`
	if string(b) != expected {
		t.Errorf("wrong result: %s", cmp.Diff(expected, string(b)))
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
			var buf strings.Builder
			if err := addAttributes(&buf, test.val, test.attrs, 0); err != nil {
				t.Errorf("unexpected error")
			}
			if buf.String() != test.expected {
				t.Errorf("wrong result: %s", cmp.Diff(test.expected, buf.String()))
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
		var buf strings.Builder
		addBlocks(&buf, val, schema.NestedBlocks, 0)

		expected := `root_block_device {
  volume_type = "foo"
}
`

		if !cmp.Equal(buf.String(), expected) {
			t.Errorf("wrong output:\n%s", cmp.Diff(expected, buf.String()))
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
		var buf strings.Builder
		addBlocks(&buf, val, schema.NestedBlocks, 0)

		expected := `root_block_device {
  volume_type = "foo"
}
root_block_device {
  volume_type = "bar"
}
`

		if !cmp.Equal(buf.String(), expected) {
			t.Fatalf("wrong output:\n%s", cmp.Diff(expected, buf.String()))
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
		var buf strings.Builder
		addBlocks(&buf, val, schema.NestedBlocks, 0)

		expected := `root_block_device {
  volume_type = "bar"
}
root_block_device {
  volume_type = "foo"
}
`

		if !cmp.Equal(buf.String(), expected) {
			t.Fatalf("wrong output:\n%s", cmp.Diff(expected, buf.String()))
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
		var buf strings.Builder
		addBlocks(&buf, val, schema.NestedBlocks, 0)

		expected := `root_block_device "1" {
  volume_type = "foo"
}
root_block_device "2" {
  volume_type = "bar"
}
`

		if !cmp.Equal(buf.String(), expected) {
			t.Fatalf("wrong output:\n%s", cmp.Diff(expected, buf.String()))
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
			"foo_list": {AttributeType: cty.List(cty.Number), Optional: true},
			"foo_json": {AttributeType: cty.String, Optional: true},
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
