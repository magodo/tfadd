package tfadd

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTFAdd_init(t *testing.T) {
	cases := []struct {
		name      string
		providers []string
		expect    interface{}
	}{
		{
			name:      "no provider specified",
			providers: nil,
			expect:    regexp.MustCompile(`^$`),
		},
		{
			name:      "azurerm provider only",
			providers: []string{"hashicorp/azurerm"},
			expect: regexp.MustCompile(`^terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "\d+\.\d+\.\d+"
    }
  }
}
$`),
		},
		{
			name:      "five providers",
			providers: []string{"hashicorp/azuread", "hashicorp/azurerm", "hashicorp/google", "hashicorp/aws", "azure/azapi"},
			expect: regexp.MustCompile(`^terraform {
  required_providers {
    azuread = {
      source = "hashicorp/azuread"
      version = "\d+\.\d+\.\d+"
    }
    azurerm = {
      source = "hashicorp/azurerm"
      version = "\d+\.\d+\.\d+"
    }
    google = {
      source = "hashicorp/google"
      version = "\d+\.\d+\.\d+"
    }
    aws = {
      source = "hashicorp/aws"
      version = "\d+\.\d+\.\d+"
    }
    azapi = {
      source = "azure/azapi"
      version = "\d+\.\d+\.\d+"
    }
  }
}
$`),
		},
		{
			name:      "unsupported provider",
			providers: []string{"invalid"},
			expect:    fmt.Errorf("Unsupported provider %q\n", "invalid"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			b, err := Init(tt.providers)
			if expect, ok := tt.expect.(error); ok {
				require.EqualError(t, err, expect.Error())
				return
			}
			require.NoError(t, err)
			require.Regexp(t, tt.expect, string(b))
		})
	}
}
