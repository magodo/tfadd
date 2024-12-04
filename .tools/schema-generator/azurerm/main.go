package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/provider"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
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
	resources := map[string]*schema.Resource{}
	for _, service := range provider.SupportedTypedServices() {
		for _, rs := range service.Resources() {
			wrapper := sdk.NewResourceWrapper(rs)
			rsWrapper, err := wrapper.Resource()
			if err != nil {
				return fmt.Errorf("wrapping Resource %q: %+v", rs.ResourceType(), err)
			}
			resources[rs.ResourceType()] = rsWrapper
		}
	}
	for _, service := range provider.SupportedUntypedServices() {
		for name, rs := range service.SupportedResources() {
			resources[name] = rs
		}
	}
	rschs := map[string]*tfschema.Schema{}
	for name, rs := range resources {
		rschs[name] = &tfschema.Schema{Block: tfpluginschema.FromSDKv2SchemaMap(rs.Schema)}
	}

	datasources := map[string]*schema.Resource{}
	for _, service := range provider.SupportedTypedServices() {
		for _, ds := range service.DataSources() {
			wrapper := sdk.NewDataSourceWrapper(ds)
			dsWrapper, err := wrapper.DataSource()
			if err != nil {
				return fmt.Errorf("wrapping DataSource %q: %+v", ds.ResourceType(), err)
			}
			datasources[ds.ResourceType()] = dsWrapper
		}
	}
	for _, service := range provider.SupportedUntypedServices() {
		for name, ds := range service.SupportedDataSources() {
			datasources[name] = ds
		}
	}
	dschs := map[string]*tfschema.Schema{}
	for name, ds := range datasources {
		dschs[name] = &tfschema.Schema{Block: tfpluginschema.FromSDKv2SchemaMap(ds.Schema)}
	}

	b, err := json.MarshalIndent(tfschema.ProviderSchema{ResourceSchemas: rschs, DatasourceSchemas: dschs}, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
