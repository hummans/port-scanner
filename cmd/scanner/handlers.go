package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/bndw/port-scanner/pkg/nmap"
	"github.com/bndw/port-scanner/pkg/repo"
)

// getScansHandler services "GET /scans", returning all scans for a given host.
func getScansHandler(db repo.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx  = c.Request.Context()
			host = strings.ToLower(c.Query("host"))
		)

		if host != "" {
			err := validateHostnameOrIP(host)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
		}

		scans, err := db.ListScans(ctx, host)
		if err != nil {
			log.Printf("failed to repo.ListScans for %s with err: %v", host, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, scans)
	}
}

// createScanHandler services "POST /scans". This involves:
// 1. Running a nmap port scan against the provided host
// 2. Storing the results in the database
// 3. Generating a diff between this scan and the previous scan
// 4. Returning the results
func createScanHandler(db repo.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx  = c.Request.Context()
			host = strings.ToLower(c.Query("host"))
		)

		err := validateHostnameOrIP(host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		scanResult, err := nmap.New().Scan(ctx, host)
		if err != nil {
			log.Printf("failed to nmap.Scan %s with err: %v", host, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		record := repo.Scan{ScanResult: *scanResult}
		scan, err := db.CreateScan(ctx, &record)
		if err != nil {
			log.Printf("failed to repo.CreateScan %#v with err: %v", scan, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		diff, err := comparePreviousScan(ctx, db, scan)
		if err != nil {
			log.Printf("failed to comparePreviousScan for id=%d with err: %v",
				scan.ID, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"scan": scan,
			"diff": diff,
		})
	}
}
