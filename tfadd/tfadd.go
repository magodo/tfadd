package tfadd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/magodo/tfadd/addr"
	"text/template"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfadd/tpl"
	"github.com/magodo/tfstate"
)

type runConfig struct {
	full   bool
	target *addr.ResourceAddr
}

var defaultStateConfig = runConfig{full: false}

func Run(ctx context.Context, tf *tfexec.Terraform, opts ...RunOption) ([]byte, error) {
	cfg := defaultStateConfig
	for _, o := range opts {
		if o, ok := o.(FailableOption); ok {
			if err := o.Error(); err != nil {
				return nil, fmt.Errorf("invalid option: %v", err)
			}
		}
		o.configureRun(&cfg)
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

	var errs error
	templates := []byte{}
	for _, res := range state.Values.RootModule.Resources {
		if cfg.target != nil {
			addr := addr.ResourceAddr{Type: res.Type, Name: res.Name}
			if *cfg.target != addr {
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
		templates = append(templates, b...)
	}

	return templates, errs
}

func Setup(providers []string) ([]byte, error) {
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
