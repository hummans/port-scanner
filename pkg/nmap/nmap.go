// package nmap defines and implements a nmap scanner interface.
package nmap

import (
	"context"
	"fmt"
	"strings"

	_nmap "github.com/Ullaakut/nmap/v2"
)

// Scanner defines an interface for running nmap port scans against a host.
// The specific ports scanned are left to the implemention.
type Scanner interface {
	Scan(ctx context.Context, host string) (*ScanResult, error)
}

// New returns a ready-to-use Scanner.
func New() Scanner {
	return &scanner{}
}

// scanner implements the Scanner interface.
type scanner struct {
}

// Scan runs a nmap scan against host for ports 1-1000.
// Scan assumes the host has already been validated.
func (s *scanner) Scan(ctx context.Context, host string) (*ScanResult, error) {
	client, err := _nmap.NewScanner(
		_nmap.WithContext(ctx),
		_nmap.WithTargets(host),
		_nmap.WithPorts(portsToScan()),
	)
	if err != nil {
		return nil, err
	}

	result, _, err := client.Run()
	if err != nil {
		return nil, err
	}
	if len(result.Hosts) != 1 {
		// result.Hosts should contain a single entry for the host we provided
		// when initializing the nmap scanner. If we get here then the 3rd party
		// nmap module is not working as expected.
		return nil, fmt.Errorf("unexpected nmap error")
	}

	var ports []PortStatus
	for _, port := range result.Hosts[0].Ports {
		ports = append(ports, PortStatus{
			Port:     port.ID,
			Protocol: port.Protocol,
			State:    port.State.String(),
			Name:     port.Service.Name,
		})
	}

	return &ScanResult{
		Host:  host,
		Ports: ports,
	}, nil
}

// ScanResult defines the results of a nmap scan
type ScanResult struct {
	Host  string       `json:"host"`
	Ports []PortStatus `json:"ports"`
}

// PortStatus defines the status of a single port
type PortStatus struct {
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
	State    string `json:"state"`
	Name     string `json:"name"`
}

func (p *PortStatus) String() string {
	return fmt.Sprintf("%d/%s %s", p.Port, p.Protocol, p.State)
}

// portsToScan returns a comma-delimited string of port numbers to scan,
// beginning at 1 and ending at 1000.
func portsToScan() string {
	portRange := make([]string, 1000)
	for i := range portRange {
		portRange[i] = fmt.Sprintf("%d", i+1)
	}

	return strings.Join(portRange, ",")
}
