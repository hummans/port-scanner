package repo_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bndw/port-scanner/pkg/nmap"
	"github.com/bndw/port-scanner/pkg/repo"
)

func TestNewRepo(t *testing.T) {
	const testDB = "./tmp.db"
	defer os.Remove(testDB)
	_, err := repo.New(testDB)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateScan(t *testing.T) {
	const testDB = "./tmp.db"
	defer os.Remove(testDB)

	r, err := repo.New(testDB)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	scan := repo.Scan{
		ScanResult: nmap.ScanResult{
			Host: "example.com",
			Ports: []nmap.PortStatus{
				{
					Port:     uint16(22),
					Protocol: "tcp",
					State:    "open",
					Name:     "ssh",
				},
			},
		},
	}

	result, err := r.CreateScan(context.TODO(), &scan)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, scan.Host, result.Host)
	assert.Equal(t, len(scan.Ports), len(result.Ports))
	assert.Equal(t, scan.Ports[0], result.Ports[0])
}

func TestListScans(t *testing.T) {
	const testDB = "./tmp.db"
	defer os.Remove(testDB)

	r, err := repo.New(testDB)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	scans := []repo.Scan{
		{
			ScanResult: nmap.ScanResult{
				Host: "example.com",
				Ports: []nmap.PortStatus{
					{
						Port:     uint16(22),
						Protocol: "tcp",
						State:    "open",
						Name:     "ssh",
					},
				},
			},
		},
		{
			ScanResult: nmap.ScanResult{
				Host: "example.com",
				Ports: []nmap.PortStatus{
					{
						Port:     uint16(22),
						Protocol: "tcp",
						State:    "closed",
						Name:     "ssh",
					},
				},
			},
		},
	}

	// Insert test records
	for _, scan := range scans {
		_, err := r.CreateScan(context.TODO(), &scan)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Set the expected IDs on the inserted records so we can assert struct
	// equality.
	scans[0].ID = int64(1)
	scans[1].ID = int64(2)

	// Ensure we can fetch them
	results, err := r.ListScans(context.TODO(), "example.com")
	assert.Nil(t, err)
	assert.Equal(t, len(scans), len(results))
	assert.Equal(t, int64(1), results[0].ID)
	assert.Equal(t, int64(2), results[1].ID)
	assert.Equal(t, "example.com", results[0].Host)
	assert.Equal(t, "example.com", results[1].Host)
	assert.Equal(t, len(scans[0].Ports), len(results[0].Ports))
	assert.Equal(t, len(scans[1].Ports), len(results[1].Ports))
	assert.Equal(t, scans[0].Ports, results[0].Ports)
	assert.Equal(t, scans[1].Ports, results[1].Ports)
}
