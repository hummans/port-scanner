package main

import (
	"testing"
)

func TestValidateHostnameOrIP(t *testing.T) {
	tests := []struct {
		Name  string
		Input string
		Fail  bool
	}{
		{"valid_ip", "10.0.0.1", false},
		{"valid_ip_trailing_dot", "10.0.0.1.", false},
		{"invalid_ip", "10.0.0", true},
		{"invalid_ip_2", "10", true},
		{"valid_host", "example.com", false},
		{"valid_host_trailing_dot", "example.com.", false},
		{"valid_host_subdomain", "x.example.com", false},
		{"valid_host_fun_tld", "example.lol", false},
		{"invalid_host", "example", true},
		{"invalid_host_scheme_http", "http://example", true},
		{"invalid_host_scheme_https", "https://example", true},
		{"invalid_host_non_ascii", "ⒷⒶⒹ", true},
		{"invalid_host_short_label", "foo..bar", true},
		{"invalid_host_long_label", "foo.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.bar", true},
		{"invalid_host_tld_only_numbers", "foo.123", true},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := validateHostnameOrIP(tt.Input)
			if !tt.Fail && err != nil {
				t.Error(err)
			}
		})
	}
}
