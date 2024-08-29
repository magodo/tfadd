package main

import (
	"bytes"
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

type arrayFlag []string

func (l *arrayFlag) String() string {
	return "array flag"
}

func (l *arrayFlag) Set(value string) error {
	*l = append(*l, value)
	return nil
}

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
  -mask-sensitive     Mask sensitive properties
  -target=addr        Only generate for the specified resource (can specify multiple times)
  -chdir              Change the current working directory
`
	return strings.TrimSpace(helpText)
}

func (r *stateCommand) Run(args []string) int {
	fset := defaultFlagSet("state")
	flagFull := fset.Bool("full", false, "Whether to generate all non-computed properties")
	flagChdir := fset.String("chdir", ".", "Change the current working directory")
	flagMaskSensitive := fset.Bool("mask-sensitive", false, "Whether to mask sensitive properties")
	var flagTargets arrayFlag
	fset.Var(&flagTargets, "target", "Only generate for the specified resource")
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
	tf, err := tfexec.NewTerraform(*flagChdir, execPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	opts := []tfadd.OptionSetter{
		tfadd.Full(*flagFull),
		tfadd.MaskSenstitive(*flagMaskSensitive),
	}

	var template []byte
	if len(flagTargets) == 0 {
		b, err := tfadd.State(ctx, tf, opts...)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return 1
		}
		template = b
	} else {
		bs, err := tfadd.StateForTargets(ctx, tf, flagTargets, opts...)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return 1
		}
		template = bytes.Join(bs, nil)
	}
	fmt.Println(string(template))
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
