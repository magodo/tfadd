package tfadd

import (
	"github.com/magodo/tfadd/providers/aws"
	"github.com/magodo/tfadd/providers/azurerm"
	"github.com/magodo/tfadd/providers/google"
	"github.com/magodo/tfadd/schema"
)

var sdkProviderSchemas = map[string]schema.ProviderSchema{
	"registry.terraform.io/hashicorp/azurerm": azurerm.ProviderSchemaInfo,
	"registry.terraform.io/hashicorp/aws":     aws.ProviderSchemaInfo,
	"registry.terraform.io/hashicorp/google":  google.ProviderSchemaInfo,
}
