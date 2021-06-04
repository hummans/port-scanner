package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bndw/port-scanner/pkg/nmap"
	"github.com/bndw/port-scanner/pkg/repo"
)

func TestComparePreviousScan(t *testing.T) {
	ctx := context.TODO()

	// Setup a repo
	const testDB = "./tmp.db"
	defer os.Remove(testDB)
	r, err := repo.New(testDB)
	if err != nil {
		t.Fatal(err)
	}

	// Insert the first scan
	_, err = r.CreateScan(ctx, &repo.Scan{
		ScanResult: nmap.ScanResult{
			Host: "example.com",
			Ports: []nmap.PortStatus{
				{
					Port:     uint16(22),
					Protocol: "tcp",
					State:    "open",
					Name:     "ssh",
				},
				{
					Port:     uint16(80),
					Protocol: "tcp",
					State:    "open",
					Name:     "http",
				},
				{
					Port:     uint16(666),
					Protocol: "tcp",
					State:    "closed",
					Name:     "http",
				},
			},
		},
	})
	assert.Nil(t, err)

	// Insert the second scan
	// changes:
	// - 22 removed
	// - 443 added
	// - 666 state changed from closed => open
	scan, err := r.CreateScan(ctx, &repo.Scan{
		ScanResult: nmap.ScanResult{
			Host: "example.com",
			Ports: []nmap.PortStatus{
				{
					Port:     uint16(80),
					Protocol: "tcp",
					State:    "open",
					Name:     "http",
				},
				{
					Port:     uint16(443),
					Protocol: "tcp",
					State:    "open",
					Name:     "https",
				},
				{
					Port:     uint16(666),
					Protocol: "tcp",
					State:    "open",
					Name:     "http",
				},
			},
		},
	})
	assert.Nil(t, err)

	diff, err := comparePreviousScan(ctx, r, scan)
	assert.Nil(t, err)

	expected := []portDiff{
		{From: "22/tcp open", To: ""},
		{From: "666/tcp closed", To: "666/tcp open"},
		{From: "", To: "443/tcp open"},
	}

	assert.Equal(t, expected, diff)
}
