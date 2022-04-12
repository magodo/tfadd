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
			providers: []string{"azurerm"},
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
			name:      "three providers only",
			providers: []string{"azurerm", "google", "aws"},
			expect: regexp.MustCompile(`^terraform {
  required_providers {
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
