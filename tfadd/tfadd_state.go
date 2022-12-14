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
	// Whether the generated config contains all the non-computed properties?
	// Set via Full option.
	full bool

	// Only generate for the specified target address.
	// Set via Target option.
	target *addr.ResourceAddr
}

func defaultStateConfig() stateConfig {
	return stateConfig{
		full: false,
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

	gen := func(pschs *tfjson.ProviderSchemas, res tfstate.StateResource, full bool) ([]byte, error) {
		if res.Mode != tfjson.ManagedResourceMode {
			return nil, nil
		}
		psch, ok := pschs.Schemas[res.ProviderName]
		if !ok {
			return nil, fmt.Errorf("no provider named %s found in provider schemas of current workspace", res.ProviderName)
		}
		rsch, ok := psch.ResourceSchemas[res.Type]
		if !ok {
			return nil, fmt.Errorf("no resource type %s found in provider's schema", res.Type)
		}
		b, err := internal.StateToTpl(&res, rsch.Block)
		if err != nil {
			return nil, fmt.Errorf("generate template from state for %s: %v", res.Type, err)
		}
		if !full {
			sdkPsch, ok := sdkProviderSchemas[res.ProviderName]
			if !ok {
				return b, nil
			}
			sch, ok := sdkPsch.ResourceSchemas[res.Type]
			if !ok {
				return b, nil
			}
			b, err = internal.TuneTpl(*sch, b, res.Type)
			if err != nil {
				return nil, fmt.Errorf("tune template for %s: %v", res.Type, err)
			}
		}
		return b, nil
	}

	if cfg.target == nil {
		var templates []byte
		var errs error
		var genForModule func(pschs *tfjson.ProviderSchemas, module tfstate.StateModule, full bool)
		genForModule = func(pschs *tfjson.ProviderSchemas, module tfstate.StateModule, full bool) {
			if module.Address != "" {
				templates = append(templates, []byte("# "+module.Address+"\n")...)
			}
			for _, res := range module.Resources {
				b, err := gen(pschs, *res, cfg.full)
				if err != nil {
					errs = multierror.Append(errs, err)
					continue
				}
				if b == nil {
					continue
				}
				templates = append(templates, b...)
			}
			for _, mod := range module.ChildModules {
				genForModule(pschs, *mod, full)
			}
		}
		genForModule(pschs, *state.Values.RootModule, cfg.full)
		return templates, errs
	}

	module := state.Values.RootModule
	for i := 0; i < len(cfg.target.ModuleAddr); i++ {
		moduleAddr := addr.ModuleAddr(cfg.target.ModuleAddr[:i+1]).String()
		var found bool
		for _, cm := range module.ChildModules {
			if cm.Address == moduleAddr {
				module = cm
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("failed to find module %s", moduleAddr)
		}
	}

	var targetResource *tfstate.StateResource
	for _, res := range module.Resources {
		if res.Type != cfg.target.Type || res.Name != cfg.target.Name {
			continue
		}
		targetResource = res
		break
	}
	if targetResource == nil {
		return nil, fmt.Errorf("can't find target resource")
	}
	return gen(pschs, *targetResource, cfg.full)
}
