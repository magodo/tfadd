package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-provider-google/google/provider"
	tfschema "github.com/magodo/tfadd/schema"
	"github.com/magodo/tfpluginschema"
)

func main() {
	provider := provider.Provider()
	rschs := map[string]*tfschema.Schema{}
	for name, rs := range provider.ResourcesMap {
		rschs[name] = &tfschema.Schema{Block: tfpluginschema.FromSDKv2SchemaMap(rs.Schema)}
	}
	dschs := map[string]*tfschema.Schema{}
	for name, ds := range provider.DataSourcesMap {
		dschs[name] = &tfschema.Schema{Block: tfpluginschema.FromSDKv2SchemaMap(ds.Schema)}
	}
	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: rschs, DatasourceSchemas: dschs}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
