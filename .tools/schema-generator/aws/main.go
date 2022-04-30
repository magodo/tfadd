package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tfschema "github.com/magodo/tfadd/schema/legacy"
)

func main() {
	schemas := map[string]*tfschema.Schema{}
	for name, rs := range provider.Provider().ResourcesMap {
		schemas[name] = &tfschema.Schema{Block: tfschema.FromProviderSchemaMap(rs.Schema)}
	}
	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: schemas}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
