package main

import (
	"context"
	"fmt"
	"log"

	"github.com/r3labs/diff/v2"

	"github.com/bndw/port-scanner/pkg/repo"
)

// portDiff captures changes to a port between 2 scans.
type portDiff struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// comparePreviousScan takes the provided ScanResult and generates a diff
// against the previous scan of the same host. If no previous scan exists
// an empty diff is returned.
func comparePreviousScan(ctx context.Context, db repo.Repo, currentScan *repo.Scan) ([]portDiff, error) {
	if currentScan.ID <= 1 {
		return []portDiff{}, nil
	}
	previousID := currentScan.ID - 1

	hostScans, err := db.ListScans(ctx, currentScan.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to ListScans for host=%s with err: %v",
			currentScan.Host, err)
	}

	var previousScan *repo.Scan
	for _, scan := range hostScans {
		// ListScans is ordered by CreatedAt so we can safely grab the first scan
		// after currentScan.
		if scan.ID < currentScan.ID {
			previousScan = &scan
			break
		}
	}
	if previousScan == nil {
		log.Printf("no previous scans for host=%s. ListScans return %d",
			currentScan.Host, len(hostScans))
		return []portDiff{}, nil
	}

	// String ports for easy diffing
	var from []string
	for _, p := range previousScan.Ports {
		from = append(from, p.String())
	}
	var to []string
	for _, p := range currentScan.Ports {
		to = append(to, p.String())
	}

	// Generate a diff with a 3rd party lib
	changeLog, err := diff.Diff(from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to diff versions from=%d to=%d with err: %v",
			previousID, currentScan.ID, err)
	}

	// Convert the 3rd party ChangeLog to []portDiff
	diff := make([]portDiff, 0)
	for _, change := range changeLog {
		diff = append(diff, portDiff{
			From: changeString(change.From),
			To:   changeString(change.To),
		})
	}

	return diff, nil
}

// changeString is a util function for getting the string type back from
// the diff.Diff results.
func changeString(x interface{}) string {
	switch x.(type) {
	case string:
		return x.(string)
	default:
		// Always type=nil, but this way we don't have to add a return at the end.
		return ""
	}
}
