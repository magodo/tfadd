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
	rschs := map[string]*tfschema.Schema{}
	for name, rs := range sch.ResourceSchemas {
		rschs[name] = &tfschema.Schema{Block: rs.Block}
	}

	dschs := map[string]*tfschema.Schema{}
	for name, ds := range sch.DataSourceSchemas {
		dschs[name] = &tfschema.Schema{Block: ds.Block}
	}

	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: rschs, DatasourceSchemas: dschs}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
