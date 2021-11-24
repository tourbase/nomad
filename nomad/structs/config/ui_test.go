package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUIConfig_Merge(t *testing.T) {

	fullConfig := &UIConfig{
		Enabled: true,
		Consul: &ConsulUIConfig{
			BaseURL: "http://consul.example.com:8500",
		},
		Vault: &VaultUIConfig{
			BaseURL: "http://vault.example.com:8200",
		},
	}

	testCases := []struct {
		name   string
		left   *UIConfig
		right  *UIConfig
		expect *UIConfig
	}{
		{
			name:   "merge onto empty config",
			left:   &UIConfig{},
			right:  fullConfig,
			expect: fullConfig,
		},
		{
			name:   "merge in a nil config",
			left:   fullConfig,
			right:  nil,
			expect: fullConfig,
		},
		{
			name: "merge onto zero-values",
			left: &UIConfig{
				Enabled: false,
				Consul: &ConsulUIConfig{
					BaseURL: "http://consul-other.example.com:8500",
				},
			},
			right:  fullConfig,
			expect: fullConfig,
		},
		{
			name: "merge from zero-values",
			left: &UIConfig{
				Enabled: true,
				Consul: &ConsulUIConfig{
					BaseURL: "http://consul-other.example.com:8500",
				},
			},
			right: &UIConfig{},
			expect: &UIConfig{
				Enabled: false,
				Consul: &ConsulUIConfig{
					BaseURL: "http://consul-other.example.com:8500",
				},
				Vault: &VaultUIConfig{},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := tc.left.Merge(tc.right)
			require.Equal(t, tc.expect, result)
		})
	}

}
