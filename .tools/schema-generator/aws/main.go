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
	provider, err := provider.New(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
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
