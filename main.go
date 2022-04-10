package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfadd/providers/aws"
	"github.com/magodo/tfadd/providers/azurerm"
	"github.com/magodo/tfadd/schema/legacy"
	"github.com/magodo/tfadd/tpl"
	"github.com/magodo/tfstate"
)

var sdkProviderSchemas = map[string]legacy.ProviderSchema{
	"registry.terraform.io/hashicorp/azurerm": azurerm.ProviderSchemaInfo,
	"registry.terraform.io/hashicorp/aws":     aws.ProviderSchemaInfo,
}

func main() {
	log.SetFlags(0)
	ctx := context.TODO()
	av := fs.AnyVersion{
		Product: &product.Terraform,
	}
	execPath, err := av.Find(ctx)
	if err != nil {
		log.Fatal(err)
	}
	tf, err := tfexec.NewTerraform(".", execPath)
	if err != nil {
		log.Fatal(err)
	}
	rawState, err := tf.Show(ctx)
	if err != nil {
		log.Fatalf("show state: %v", err)
	}
	if rawState == nil || rawState.Values == nil {
		log.Fatalf("No state")
	}
	pschs, err := tf.ProvidersSchema(ctx)
	if err != nil {
		log.Fatalf("get provider schemas: %v", err)
	}
	state, err := tfstate.FromJSONState(rawState, pschs)
	if err != nil {
		log.Fatal(fmt.Sprintf("from json state: %v", err))
	}

	var errs error
	templates := []byte{}
	for _, res := range state.Values.RootModule.Resources {
		if res.Mode != tfjson.ManagedResourceMode {
			log.Printf("\tSkipping %s, since it is not a managed resource\n", res.Address)
			continue
		}
		psch, ok := pschs.Schemas[res.ProviderName]
		if !ok {
			log.Printf("\tSkipping %s, since can't find the provider schema for %s\n", res.Address, res.ProviderName)
			continue
		}
		rsch, ok := psch.ResourceSchemas[res.Type]
		if !ok {
			log.Printf("\tSkipping %s, since can't find the resource schema in the provider schema\n", res.Address)
			continue
		}
		b, err := tpl.StateToTpl(res, rsch.Block)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("generate template from state for %s: %v", res.Type, err))
		}
		sdkPsch, ok := sdkProviderSchemas[res.ProviderName]
		if !ok {
			log.Printf("\tSkipping %s, since can't find the resource schema in the SDK provider schema\n", res.Address)
			continue
		}
		b, err = sdkPsch.TuneTpl(b, res.Type)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("tune template for %s: %v", res.Type, err))
		}
		templates = append(templates, b...)
	}

	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Println(string(templates))
}
