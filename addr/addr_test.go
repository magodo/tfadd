package addr

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/stretchr/testify/require"
)

func ptr[T any](in T) *T {
	return &in
}

func TestParseModuleAddr(t *testing.T) {
	cases := []struct {
		name  string
		input string
		addr  ModuleAddr
		err   bool
	}{
		{
			name:  "one module",
			input: "module.mod1",
			addr: []ModuleStep{
				{
					Name: "mod1",
				},
			},
		},
		{
			name:  "module instance (key)",
			input: `module.mod1["foo"]`,
			addr: []ModuleStep{
				{
					Name: "mod1",
					Key:  ptr("foo"),
				},
			},
		},
		{
			name:  "module instance (idx)",
			input: `module.mod1[0]`,
			addr: []ModuleStep{
				{
					Name:  "mod1",
					Index: ptr(0),
				},
			},
		},
		{
			name:  "nested module instance",
			input: `module.mod1[0].module.mod2["foo"].module.mod3`,
			addr: []ModuleStep{
				{
					Name:  "mod1",
					Index: ptr(0),
				},
				{
					Name: "mod2",
					Key:  ptr("foo"),
				},
				{
					Name: "mod3",
				},
			},
		},
		{
			name:  "invalid module",
			input: "mod1",
			err:   true,
		},
		{
			name:  "invalid module instance",
			input: "module.mod1[]",
			err:   true,
		},
		{
			name:  "invalid module instance key",
			input: "module.mod1[xyz]",
			err:   true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := ParseModuleAddr(tt.input)
			if tt.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.addr, addr)
		})
	}
}

func TestParseResourceAddr(t *testing.T) {
	cases := []struct {
		name  string
		input string
		addr  ResourceAddr
		err   bool
	}{
		{
			name:  "resource only",
			input: "null_resource.test",
			addr: ResourceAddr{
				Mode: tfjson.ManagedResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "data source only",
			input: "data.null_resource.test",
			addr: ResourceAddr{
				Mode: tfjson.DataResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "resource with module",
			input: "module.mod1.null_resource.test",
			addr: ResourceAddr{
				ModuleAddr: []ModuleStep{
					{
						Name: "mod1",
					},
				},
				Mode: tfjson.ManagedResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "data source with module",
			input: "module.mod1.data.null_resource.test",
			addr: ResourceAddr{
				ModuleAddr: []ModuleStep{
					{
						Name: "mod1",
					},
				},
				Mode: tfjson.DataResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "resource with module instance (key)",
			input: `module.mod1["foo"].null_resource.test`,
			addr: ResourceAddr{
				ModuleAddr: []ModuleStep{
					{
						Name: "mod1",
						Key:  ptr("foo"),
					},
				},
				Mode: tfjson.ManagedResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "resource with module instance (idx)",
			input: `module.mod1[0].null_resource.test`,
			addr: ResourceAddr{
				ModuleAddr: []ModuleStep{
					{
						Name:  "mod1",
						Index: ptr(0),
					},
				},
				Mode: tfjson.ManagedResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "resource with nested module instance",
			input: `module.mod1[0].module.mod2["foo"].module.mod3.null_resource.test`,
			addr: ResourceAddr{
				ModuleAddr: []ModuleStep{
					{
						Name:  "mod1",
						Index: ptr(0),
					},
					{
						Name: "mod2",
						Key:  ptr("foo"),
					},
					{
						Name: "mod3",
					},
				},
				Mode: tfjson.ManagedResourceMode,
				Type: "null_resource",
				Name: "test",
			},
		},
		{
			name:  "invalid resource addr",
			input: "null_resource",
			err:   true,
		},
		{
			name:  "invalid resource addr with module",
			input: "mod1.null_resource.test",
			err:   true,
		},
		{
			name:  "invalid resource addr with module instance",
			input: "module.mod1[].null_resource.test",
			err:   true,
		},
		{
			name:  "invalid resource addr with module instance key",
			input: "module.mod1[xyz].null_resource.test",
			err:   true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := ParseResourceAddr(tt.input)
			if tt.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.addr, *addr)
		})
	}
}
