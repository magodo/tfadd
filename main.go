package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/magodo/tfadd/tfadd"
	"github.com/mitchellh/cli"
)

func defaultFlagSet(name string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)
	f.Usage = func() {}
	return f
}

type initCommand struct{}

func (s *initCommand) Help() string {
	helpText := `
Usage: tfadd [global options] init [options] [providers]

  Generate Terraform setting that pins the provider versions to standard output.
`
	return strings.TrimSpace(helpText)
}

func (s *initCommand) Run(args []string) int {
	b, err := tfadd.Init(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	if len(b) == 0 {
		return 0
	}
	fmt.Fprintln(os.Stdout, string(b))
	return 0
}

func (s *initCommand) Synopsis() string {
	return "Setup the Terraform setting"
}

type stateCommand struct{}

func (r *stateCommand) Help() string {
	helpText := `
Usage: tfadd [global options] state [options]

  Generates resource template from Terraform state to standard output.

Options:

  -full               Output all non-computed properties in the generated config
  -target=addr        Only generate for the specified resource
`
	return strings.TrimSpace(helpText)
}

func (r *stateCommand) Run(args []string) int {
	fset := defaultFlagSet("state")
	flagFull := fset.Bool("full", false, "Whether to generate all non-computed properties")
	flagTarget := fset.String("target", "", "Only generate for the specified resource")
	if err := fset.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}

	ctx := context.TODO()
	av := fs.AnyVersion{
		Product: &product.Terraform,
	}
	execPath, err := av.Find(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	tf, err := tfexec.NewTerraform(".", execPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	opts := []tfadd.StateOption{tfadd.Full(*flagFull)}
	if *flagTarget != "" {
		opts = append(opts, tfadd.Target(*flagTarget))
	}
	templates, err := tfadd.State(ctx, tf, opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	fmt.Println(string(templates))
	return 0
}

func (r *stateCommand) Synopsis() string {
	return "Generate Terraform configuration"
}

func main() {
	c := cli.NewCLI("tfadd", "dev")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"state": func() (cli.Command, error) { return &stateCommand{}, nil },
		"init":  func() (cli.Command, error) { return &initCommand{}, nil },
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
