package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfadd/providers/aws"
	"github.com/magodo/tfadd/providers/azurerm"
	"github.com/magodo/tfadd/providers/google"
	"github.com/magodo/tfadd/schema/legacy"
	"github.com/magodo/tfadd/tpl"
	"github.com/magodo/tfstate"
	"github.com/mitchellh/cli"
)

var sdkProviderSchemas = map[string]legacy.ProviderSchema{
	"registry.terraform.io/hashicorp/azurerm": azurerm.ProviderSchemaInfo,
	"registry.terraform.io/hashicorp/aws":     aws.ProviderSchemaInfo,
	"registry.terraform.io/hashicorp/google":  google.ProviderSchemaInfo,
}

func defaultFlagSet(name string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)
	f.Usage = func() {}
	return f
}

type setupCommand struct{}

func (s *setupCommand) Help() string {
	helpText := `
Usage: tfadd [global options] setup [options] [providers]

  Generate Terraform setting that pins the provider versions to standard output.
`
	return strings.TrimSpace(helpText)
}

func (s *setupCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No provider specified")
		return 0
	}

	type info struct {
		name    string
		source  string
		version string
	}

	var infos []info
	for _, p := range args {
		pinfo, ok := sdkProviderSchemas["registry.terraform.io/hashicorp/"+p]
		if !ok {
			fmt.Fprintf(os.Stderr, "Unsupported provider %q\n", p)
			return 1
		}
		infos = append(infos, info{
			name: p,
			source: "hashicorp/" + p,
			version: "v" // TODO,
		})

	}
}

func (s *setupCommand) Synopsis() string {
	return "Setup the Terraform setting"
}

type runCommand struct{}

func (r *runCommand) Help() string {
	helpText := `
Usage: tfadd [global options] run [options]

  Generates resource template from Terraform state to standard output.

Options:

  -full         	  Output all non-computed properties in the generated config.
`
	return strings.TrimSpace(helpText)
}

func (r *runCommand) Run(args []string) int {
	fset := defaultFlagSet("run")
	flagFull := fset.Bool("full", false, "Whether to generate all non-computed properties")
	fset.Parse(args)

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
		if !*flagFull {
			sdkPsch, ok := sdkProviderSchemas[res.ProviderName]
			if !ok {
				log.Printf("\tSkipping %s, since can't find the resource schema in the SDK provider schema\n", res.Address)
				continue
			}
			b, err = sdkPsch.TuneTpl(b, res.Type)
			if err != nil {
				errs = multierror.Append(errs, fmt.Errorf("tune template for %s: %v", res.Type, err))
			}
		}
		templates = append(templates, b...)
	}

	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Println(string(templates))
	return 0
}

func (r *runCommand) Synopsis() string {
	return "Generate Terraform configuration"
}

func main() {
	c := cli.NewCLI("tfadd", "dev")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"run":   func() (cli.Command, error) { return &runCommand{}, nil },
		"setup": func() (cli.Command, error) { return &setupCommand{}, nil },
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
