package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-provider-azuread/internal/provider"
	tfschema "github.com/magodo/tfadd/schema"
	"github.com/magodo/tfpluginschema"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	provider := provider.AzureADProvider()
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
		return err
	}
	fmt.Println(string(b))
	return nil
}
