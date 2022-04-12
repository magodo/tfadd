package tfadd

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/magodo/tfadd/addr"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfadd/tpl"
	"github.com/magodo/tfstate"
)

type stateConfig struct {
	// Whether the generated config contains all the non-computed properties?
	// Set via Full option.
	full bool

	// Only generate for the specified one or more target addresses.
	// Set via Target option.
	targets   []addr.ResourceAddr
	targetMap map[addr.ResourceAddr]bool
}

var defaultStateConfig = stateConfig{
	full: false,

	targets:   []addr.ResourceAddr{},
	targetMap: map[addr.ResourceAddr]bool{},
}

func State(ctx context.Context, tf *tfexec.Terraform, opts ...StateOption) ([]byte, error) {
	cfg := defaultStateConfig
	for _, o := range opts {
		o.configureState(&cfg)
	}

	rawState, err := tf.Show(ctx)
	if err != nil {
		return nil, fmt.Errorf("show state: %v", err)
	}
	if rawState == nil || rawState.Values == nil {
		return nil, fmt.Errorf("no state")
	}
	pschs, err := tf.ProvidersSchema(ctx)
	if err != nil {
		return nil, fmt.Errorf("get provider schemas: %v", err)
	}
	state, err := tfstate.FromJSONState(rawState, pschs)
	if err != nil {
		return nil, fmt.Errorf("from json state: %v", err)
	}

	// templateMap is only used when -target is specified.
	// It is mainly used caching the template and later sort it to the same order as option order in CLI.
	templateMap := map[addr.ResourceAddr][]byte{}
	hasTarget := len(cfg.targets) != 0

	var errs error
	templates := []byte{}

	for _, res := range state.Values.RootModule.Resources {
		raddr := addr.ResourceAddr{Type: res.Type, Name: res.Name}
		if hasTarget {
			if !cfg.targetMap[raddr] {
				continue
			}
		}
		if res.Mode != tfjson.ManagedResourceMode {
			continue
		}
		psch, ok := pschs.Schemas[res.ProviderName]
		if !ok {
			continue
		}
		rsch, ok := psch.ResourceSchemas[res.Type]
		if !ok {
			continue
		}
		b, err := tpl.StateToTpl(res, rsch.Block)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("generate template from state for %s: %v", res.Type, err))
		}
		if !cfg.full {
			sdkPsch, ok := sdkProviderSchemas[res.ProviderName]
			if !ok {
				continue
			}
			b, err = sdkPsch.TuneTpl(b, res.Type)
			if err != nil {
				errs = multierror.Append(errs, fmt.Errorf("tune template for %s: %v", res.Type, err))
			}
		}
		if hasTarget {
			templateMap[raddr] = b
		} else {
			templates = append(templates, b...)
		}
	}

	if hasTarget {
		for _, raddr := range cfg.targets {
			templates = append(templates, templateMap[raddr]...)
		}
	}

	return templates, errs
}

func Init(providers []string) ([]byte, error) {
	if len(providers) == 0 {
		return nil, nil
	}

	type info struct {
		Name    string
		Source  string
		Version string
	}

	var infos []info
	for _, p := range providers {
		pinfo, ok := sdkProviderSchemas["registry.terraform.io/hashicorp/"+p]
		if !ok {
			return nil, fmt.Errorf("Unsupported provider %q\n", p)
		}
		infos = append(infos, info{
			Name:    p,
			Source:  "hashicorp/" + p,
			Version: pinfo.Version,
		})
	}

	out := bytes.Buffer{}
	if err := template.Must(template.New("setup").Parse(`terraform {
  required_providers {
  {{- range . }}
	{{.Name}} = {
	  source = "hashicorp/{{.Name}}"
	  version = "{{.Version}}"
	}
  {{- end }}
  }
}
`)).Execute(&out, infos); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
