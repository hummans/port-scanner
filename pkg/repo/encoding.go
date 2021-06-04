package repo

import (
	"encoding/json"
	"time"

	"github.com/bndw/port-scanner/pkg/nmap"
)

// encodeScanResults encodes a ScanResult.Ports into a []byte so it
// can be stored as a blob in the database.
// PERF: JSON encoding may be swapped with a more performant encoding.
func encodeScanResults(ps []nmap.PortStatus) ([]byte, error) {
	return json.Marshal(ps)
}

// decodeScanResults decodes a JSON-encoded ScanResult.Ports.
// PERF: JSON encoding may be swapped with a more performant encoding.
func decodeScanResults(results []byte) ([]nmap.PortStatus, error) {
	var ports []nmap.PortStatus
	err := json.Unmarshal(results, &ports)
	return ports, err
}

// decodeCreatedAt decodes a database datetime string into a time.Time.
func decodeCreatedAt(ts string) (time.Time, error) {
	return time.Parse(time.RFC3339, ts)
}
