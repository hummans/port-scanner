package nmap

import (
	"context"
	"strings"
	"testing"
)

func TestNmapScan(t *testing.T) {
	tests := []struct {
		Name string
		Host string
	}{
		{"hostname", "example.com"},
		{"ip", "142.250.69.206"}, // dig google.com +short
	}

	ctx := context.TODO()
	client := New()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := client.Scan(ctx, tt.Host)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(result)
		})
	}
}

func TestPortsToScan(t *testing.T) {
	var (
		ports   = strings.Split(portsToScan(), ",")
		minPort = ports[0]
		maxPort = ports[len(ports)-1]
	)

	if len(ports) != 1000 {
		t.Fatal("expected 1000 ports")
	}
	if minPort != "1" {
		t.Fatal("ports to scan should start with 1")
	}
	if maxPort != "1000" {
		t.Fatal("ports to scan should end with 1000")
	}
}
