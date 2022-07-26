package tfadd

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/checkpoint"
	tffs "github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/src"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/stretchr/testify/require"
)

const (
	// Controls whether to run the e2e test.
	ENV_TFADD_E2E = "TFADD_E2E"

	// Once set, ignore the `tfadd init` and `terraform init`, but you should ensure the `dev_overrides` is set properly in the .terraformrc.
	// This is mainly to avoid downloading the providers from (or even interacting with) registry, for poor souls that have bad networking (like me)...
	ENV_TFADD_DEV_PROVIDER = "TFADD_DEV_PROVIDER"
)

func TestTFAdd_state(t *testing.T) {
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
			options:   []StateOption{KeepDefaultValueAttr(false)},
			expect: `resource "azurerm_resource_group" "a" {
  id       = "/subscriptions/REDACTED/resourceGroups/foo"
  location = "eastus2"
  name     = "foo"
  tags     = null
}
resource "azurerm_resource_group" "b" {
  id       = "/subscriptions/REDACTED/resourceGroups/bar"
  location = "eastus2"
  name     = "bar"
  tags     = null
}
`,
		},
		{
			name:      "generate one resource",
			statefile: "azurerm_resource_groups",
			options:   []StateOption{Target("azurerm_resource_group.a")},
			expect: `resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
}
`,
		},
		{
			name:      "generate two resources",
			statefile: "azurerm_resource_groups",
			options: []StateOption{
				Target("azurerm_resource_group.b"),
				Target("azurerm_resource_group.a"),
			},
			expect: `resource "azurerm_resource_group" "b" {
  location = "eastus2"
  name     = "bar"
}
resource "azurerm_resource_group" "a" {
  location = "eastus2"
  name     = "foo"
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
				b, err := Init([]string{"azurerm", "google", "aws"})
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(wsp, "terraform.tf"), b, 0644))
				require.NoError(t, tf.Init(ctx))
			}

			b, err := State(ctx, tf, tt.options...)
			if tt.expectError != nil {
				require.Regexp(t, tt.expectError, err.Error())
				return
			}
			require.Equal(t, tt.expect, string(b))
		})
	}
}
