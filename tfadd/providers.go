package tfadd

import (
	"github.com/magodo/tfadd/providers/aws"
	"github.com/magodo/tfadd/providers/azurerm"
	"github.com/magodo/tfadd/providers/google"
	"github.com/magodo/tfadd/schema"
)

type providerInfo struct {
	FQName    string
	SDKSchema schema.ProviderSchema
}

var supportedProviders = map[string]providerInfo{
	"hashicorp/azurerm": {
		FQName:    "registry.terraform.io/hashicorp/azurerm",
		SDKSchema: azurerm.ProviderSchemaInfo,
	},
	"hashicorp/aws": {
		FQName:    "registry.terraform.io/hashicorp/aws",
		SDKSchema: aws.ProviderSchemaInfo,
	},
	"hashicorp/google": {
		FQName:    "registry.terraform.io/hashicorp/google",
		SDKSchema: google.ProviderSchemaInfo,
	},
}
