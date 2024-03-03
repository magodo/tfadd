package tfadd

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/checkpoint"
	tffs "github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/src"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/magodo/tfstate/terraform/jsonschema"
	"github.com/stretchr/testify/require"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

const (
	// Controls whether to run the e2e test.
	ENV_TFADD_E2E = "TFADD_E2E"

	// Once set, ignore the `tfadd init` and `terraform init`, but you should ensure the `dev_overrides` is set properly in the .terraformrc.
	// This is mainly to avoid downloading the providers from (or even interacting with) registry, for poor souls that have bad networking (like me)...
	ENV_TFADD_DEV_PROVIDER = "TFADD_DEV_PROVIDER"
)

func TestTFAdd_resource(t *testing.T) {
	if os.Getenv(ENV_TFADD_E2E) == "" {
		t.Skipf("Skipping e2e test as %q is not set", ENV_TFADD_E2E)
	}

	const testfixture string = "./testdata/tfadd_state"

	// Ensure terraform executable
	ctx := context.TODO()
	i := install.NewInstaller()
	tfexecutable, err := i.Ensure(ctx, []src.Source{
		&tffs.AnyVersion{
			Product: &product.Terraform,
		},
		&checkpoint.LatestVersion{
			Product: product.Terraform,
		},
	})
	if err != nil {
		t.Fatalf("failed to install terraform: %v", err)
	}

	cases := []struct {
		name        string
		statefile   string
		options     []StateOption
		targets     []string
		expectError *regexp.Regexp
		expect      string
	}{
		{
			name:        "no state",
			expectError: regexp.MustCompile("^no state$"),
		},
		{
			name:      "generate all supported resources in the state, with tunning",
			statefile: "azurerm_resource_groups",
			expect: `resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
}
resource "azurerm_resource_group" "b" {
  location = "eastus2"
  name     = "bar"
}
`,
		},
		{
			name:      "generate all supported resources in the state, full",
			statefile: "azurerm_resource_groups",
			options:   []StateOption{Full(true)},
			expect: `resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
}
resource "azurerm_resource_group" "b" {
  location = "eastus2"
  name     = "bar"
}
`,
		},
		{
			name:      "generate one target resource",
			statefile: "azurerm_resource_groups",
			targets: []string{
				"azurerm_resource_group.a",
			},
			expect: `resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
}
`,
		},
		{
			name:      "generate two target resources",
			statefile: "azurerm_resource_groups",
			targets: []string{
				"azurerm_resource_group.a",
				"azurerm_resource_group.b",
			},
			expect: `resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
}
resource "azurerm_resource_group" "b" {
  location = "eastus2"
  name     = "bar"
}
`,
		},
		{
			name:      "multiple providers",
			statefile: "multiple_providers",
			expect: `resource "aws_vpc" "main" {
  tags = {
    Name = "main"
  }
}
resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
}
resource "azurerm_resource_group" "b" {
  location = "eastus2"
  name     = "bar"
}
`,
		},
		{
			name:      "module",
			statefile: "module",
			expect: `# module.mod1[0]
resource "null_resource" "test" {
}
# module.mod1[1]
resource "null_resource" "test" {
}
`,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			wsp := t.TempDir()

			if tt.statefile != "" {
				b, err := os.ReadFile(filepath.Join(testfixture, tt.statefile))
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(wsp, "terraform.tfstate"), b, 0644))
			}

			tf, err := tfexec.NewTerraform(wsp, tfexecutable)
			require.NoError(t, err)

			ctx := context.Background()
			if os.Getenv(ENV_TFADD_DEV_PROVIDER) == "" {
				b, err := Init([]string{"hashicorp/azurerm", "hashicorp/google", "hashicorp/aws", "azure/azapi"})
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(wsp, "terraform.tf"), b, 0644))
				require.NoError(t, tf.Init(ctx))
			}

			if len(tt.targets) == 0 {
				b, err := State(ctx, tf, tt.options...)
				if tt.expectError != nil {
					require.Regexp(t, tt.expectError, err.Error())
					return
				}
				require.NoError(t, err)
				require.Equal(t, tt.expect, string(b))
				return
			}

			bs, err := StateForTargets(ctx, tf, tt.targets)
			if tt.expectError != nil {
				require.Regexp(t, tt.expectError, err.Error())
				return
			}
			require.NoError(t, err)
			b := bytes.Join(bs, nil)
			require.Equal(t, tt.expect, string(b))
		})
	}
}

func TestTFAdd_provider(t *testing.T) {
	if os.Getenv(ENV_TFADD_E2E) == "" {
		t.Skipf("Skipping e2e test as %q is not set", ENV_TFADD_E2E)
	}

	// Ensure terraform executable
	ctx := context.TODO()
	i := install.NewInstaller()
	tfexecutable, err := i.Ensure(ctx, []src.Source{
		&tffs.AnyVersion{
			Product: &product.Terraform,
		},
		&checkpoint.LatestVersion{
			Product: product.Terraform,
		},
	})
	if err != nil {
		t.Fatalf("failed to install terraform: %v", err)
	}

	cases := []struct {
		name         string
		providerFQDN string
		v            map[string]interface{}
		expectError  *regexp.Regexp
		expect       string
	}{
		{
			name:         "generate provider config",
			providerFQDN: "registry.terraform.io/hashicorp/azurerm",
			v: map[string]interface{}{
				"features":  []interface{}{map[string]interface{}{}},
				"client_id": "client1",
			},
			expect: `provider "azurerm" {
  client_id = "client1"
  features {
  }
}
`,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			wsp := t.TempDir()
			tf, err := tfexec.NewTerraform(wsp, tfexecutable)
			require.NoError(t, err)
			ctx := context.Background()
			require.NoError(t, os.WriteFile(filepath.Join(wsp, "terraform.tf"), []byte(`provider "azurerm" {
  features {}
}`), 0644))
			require.NoError(t, tf.Init(ctx))

			pschs, err := tf.ProvidersSchema(context.TODO())
			require.NoError(t, err)
			psch := pschs.Schemas[tt.providerFQDN].ConfigSchema

			b, err := json.Marshal(tt.v)
			require.NoError(t, err)

			schema := jsonschema.SchemaBlockImpliedType(psch.Block)
			ctyval, err := ctyjson.Unmarshal(b, schema)

			providerName := strings.Split(tt.providerFQDN, "/")[2]
			b, err = GenerateForProvider(providerName, psch, ctyval)
			if tt.expectError != nil {
				require.Regexp(t, tt.expectError, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expect, string(b))
			return
		})
	}
}
