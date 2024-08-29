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
	sch, err := tfpluginschema.FromFWProvider(&provider.Provider{})
	if err != nil {
		log.Fatal(err)
	}
	schemas := map[string]*tfschema.Schema{}
	for name, rs := range sch.ResourceSchemas {
		schemas[name] = &tfschema.Schema{Block: rs.Block}
	}
	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: schemas}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
