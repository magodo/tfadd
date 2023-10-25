package tfadd

import (
	"context"
	"fmt"
	"strings"

	"github.com/magodo/tfadd/tfadd/internal"

	"github.com/magodo/tfadd/addr"

	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/magodo/tfstate"
)

type stateConfig struct {
	// Whether the generated config contains all the non-computed properties?
	// Set via Full option.
	full bool
}

func defaultStateConfig() stateConfig {
	return stateConfig{
		full: false,
	}
}

func State(ctx context.Context, tf *tfexec.Terraform, opts ...StateOption) ([]byte, error) {
	bs, err := fromState(ctx, tf, nil, opts...)
	if err != nil {
		return nil, err
	}
	return bs[0], nil
}

func StateForTargets(ctx context.Context, tf *tfexec.Terraform, targets []string, opts ...StateOption) ([][]byte, error) {
	var targetAddrs []addr.ResourceAddr
	for _, target := range targets {
		targetAddr, err := addr.ParseResourceAddr(target)
		if err != nil {
			return nil, err
		}
		targetAddrs = append(targetAddrs, *targetAddr)
	}
	return fromState(ctx, tf, targetAddrs, opts...)
}

func fromState(ctx context.Context, tf *tfexec.Terraform, targets []addr.ResourceAddr, opts ...StateOption) ([][]byte, error) {
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

	if len(targets) == 0 {
		b, err := GenerateForOneModuleModule(pschs, *state.Values.RootModule, cfg.full)
		if err != nil {
			return nil, err
		}
		return [][]byte{b}, nil
	}

	var out [][]byte
	for _, target := range targets {
		module := state.Values.RootModule
		for i := 0; i < len(target.ModuleAddr); i++ {
			moduleAddr := addr.ModuleAddr(target.ModuleAddr[:i+1]).String()
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
			if res.Type != target.Type || res.Name != target.Name {
				continue
			}
			targetResource = res
			break
		}
		if targetResource == nil {
			return nil, fmt.Errorf("can't find target resource")
		}
		psch, ok := pschs.Schemas[targetResource.ProviderName]
		if !ok {
			return nil, fmt.Errorf("no provider named %s found in provider schemas of current workspace", targetResource.ProviderName)
		}
		rsch, ok := psch.ResourceSchemas[targetResource.Type]
		if !ok {
			return nil, fmt.Errorf("no resource type %s found in provider's schema", targetResource.Type)
		}
		b, err := GenerateForOneResource(rsch, *targetResource, cfg.full)
		if err != nil {
			return nil, fmt.Errorf("generate for one resource: %v", err)
		}
		out = append(out, b)
	}
	return out, nil
}

func GenerateForOneModuleModule(pschs *tfjson.ProviderSchemas, module tfstate.StateModule, full bool) ([]byte, error) {
	var templates []byte
	if module.Address != "" {
		templates = append(templates, []byte("# "+module.Address+"\n")...)
	}
	for _, res := range module.Resources {
		psch, ok := pschs.Schemas[res.ProviderName]
		if !ok {
			return nil, fmt.Errorf("no provider named %s found in provider schemas of current workspace", res.ProviderName)
		}
		rsch, ok := psch.ResourceSchemas[res.Type]
		if !ok {
			return nil, fmt.Errorf("no resource type %s found in provider's schema", res.Type)
		}
		b, err := GenerateForOneResource(rsch, *res, full)
		if err != nil {
			return nil, err
		}
		if b == nil {
			continue
		}
		templates = append(templates, b...)
	}
	for _, mod := range module.ChildModules {
		ctemplates, err := GenerateForOneModuleModule(pschs, *mod, full)
		if err != nil {
			return nil, err
		}
		if ctemplates == nil {
			continue
		}
		templates = append(templates, ctemplates...)
	}
	return templates, nil
}

func GenerateForOneResource(rsch *tfjson.Schema, res tfstate.StateResource, full bool) ([]byte, error) {
	if res.Mode != tfjson.ManagedResourceMode {
		return nil, nil
	}
	b, err := internal.StateToTpl(&res, rsch.Block)
	if err != nil {
		return nil, fmt.Errorf("generate template from state for %s: %v", res.Type, err)
	}
	if !full {
		providerName := strings.TrimPrefix(res.ProviderName, "registry.terraform.io/")
		pinfo, ok := supportedProviders[providerName]
		if !ok {
			return b, nil
		}
		sdkPsch := pinfo.SDKSchema
		sch, ok := sdkPsch.ResourceSchemas[res.Type]
		if !ok {
			return b, nil
		}
		if providerName == "azure/azapi" {
			b, err = internal.TuneTpl(*sch, b, res.Type,
				map[string]bool{
					"name":      true,
					"parent_id": true,
					"identity":  true,
					"location":  true,
					"tags":      true,
				},
			)
		} else {
			b, err = internal.TuneTpl(*sch, b, res.Type, nil)
		}
		if err != nil {
			return nil, fmt.Errorf("tune template for %s: %v", res.Type, err)
		}
	}
	return b, nil
}
