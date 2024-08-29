package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/terraform-provider-azapi/internal/provider"
	tfschema "github.com/magodo/tfadd/schema"
	"github.com/magodo/tfpluginschema"
)

func main() {
	schemas := map[string]*tfschema.Schema{}
	for name, rs := range provider.AzureProvider().ResourcesMap {
		schemas[name] = &tfschema.Schema{Block: tfpluginschema.FromSDKv2ProviderSchemaMap(rs.Schema)}
	}
	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: schemas}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
