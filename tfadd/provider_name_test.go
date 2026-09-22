package tfadd

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestProviderShortName(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"registry.terraform.io/hashicorp/azurerm", "hashicorp/azurerm"},
		{"registry.opentofu.org/hashicorp/azurerm", "hashicorp/azurerm"},
		{"registry.terraform.io/azure/azapi", "azure/azapi"},
		{"registry.opentofu.org/azure/azapi", "azure/azapi"},
		// Any other host is left as it is, so it is not mistaken for a supported provider.
		{"example.com/hashicorp/azurerm", "example.com/hashicorp/azurerm"},
	}
	for _, c := range cases {
		require.Equal(t, c.want, providerShortName(c.in), c.in)
	}
}

// OpenTofu records the same provider as registry.opentofu.org/hashicorp/azurerm
// in state. The generated config must be tuned exactly as it is for
// registry.terraform.io/hashicorp/azurerm, rather than returned untuned.
func TestGenerateForOneResource_OpenTofuRegistryHost(t *testing.T) {
	rsch := &tfjson.Schema{
		Block: &tfjson.SchemaBlock{
			Attributes: map[string]*tfjson.SchemaAttribute{
				"id":         {AttributeType: cty.String, Computed: true},
				"name":       {AttributeType: cty.String, Required: true},
				"location":   {AttributeType: cty.String, Required: true},
				"managed_by": {AttributeType: cty.String, Optional: true},
				"tags":       {AttributeType: cty.Map(cty.String), Optional: true},
			},
		},
	}
	value := cty.ObjectVal(map[string]cty.Value{
		"id":         cty.StringVal("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example"),
		"name":       cty.StringVal("example"),
		"location":   cty.StringVal("westeurope"),
		"managed_by": cty.StringVal(""),
		"tags":       cty.NullVal(cty.Map(cty.String)),
	})
	generate := func(providerName string) string {
		b, err := GenerateForOneResource(rsch, tfstate.StateResource{
			Address:      "azurerm_resource_group.example",
			Mode:         tfjson.ManagedResourceMode,
			Type:         "azurerm_resource_group",
			Name:         "example",
			ProviderName: providerName,
			Value:        value,
		})
		require.NoError(t, err)
		return string(b)
	}

	fromTerraform := generate("registry.terraform.io/hashicorp/azurerm")
	require.NotContains(t, fromTerraform, "managed_by", "tuning removes the zero-valued optional attribute")
	require.Equal(t, fromTerraform, generate("registry.opentofu.org/hashicorp/azurerm"))
}
