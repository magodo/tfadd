package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tfschema "github.com/magodo/tfadd/schema"
	"github.com/magodo/tfpluginschema"
)

func main() {
	schemas := map[string]*tfschema.Schema{}
	provider, err := provider.New(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	for name, rs := range provider.ResourcesMap {
		schemas[name] = &tfschema.Schema{Block: tfpluginschema.FromSDKv2ProviderSchemaMap(rs.Schema)}
	}
	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: schemas}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
