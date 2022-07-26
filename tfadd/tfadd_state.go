package tfadd

import (
	"context"
	"fmt"
	"github.com/magodo/tfadd/tfadd/internal"

	"github.com/magodo/tfadd/addr"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
)

type stateConfig struct {
	tuneOption internal.TuneOption

	// Only generate for the specified one or more target addresses.
	// Set via Target option.
	targets   []addr.ResourceAddr
	targetMap map[addr.ResourceAddr]bool
}

func defaultStateConfig() stateConfig {
	return stateConfig{
		tuneOption: internal.TuneOption{
			KeepDefaultValueAttrs: false,
			IgnoreAttrConstraints: false,
		},

		targets:   []addr.ResourceAddr{},
		targetMap: map[addr.ResourceAddr]bool{},
	}
}

func State(ctx context.Context, tf *tfexec.Terraform, opts ...StateOption) ([]byte, error) {
	cfg := defaultStateConfig()
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
		b, err := internal.StateToTpl(res, rsch.Block)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("generate template from state for %s: %v", res.Type, err))
			continue
		}

		b, err = internal.TuneTpl(b)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("tune template (without schema) for %s: %v", res.Type, err))
			continue
		}

		// Further tune the template based on the provider schema.
		sdkPsch, ok := sdkProviderSchemas[res.ProviderName]
		if ok {
			sch, ok := sdkPsch.ResourceSchemas[res.Type]
			if ok {
				b, err = internal.TuneTplWithSchema(b, *sch, &cfg.tuneOption)
				if err != nil {
					errs = multierror.Append(errs, fmt.Errorf("tune template (with schema) for %s: %v", res.Type, err))
					continue
				}
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
