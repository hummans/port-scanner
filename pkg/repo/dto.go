package repo

import (
	"time"

	"github.com/bndw/port-scanner/pkg/nmap"
)

// Scan wraps a ScanResult with DB fields.
type Scan struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	nmap.ScanResult
}
