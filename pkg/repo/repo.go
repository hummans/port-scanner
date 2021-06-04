// package repo defines and implements the storage (repository) layer for Scans.
package repo

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bndw/port-scanner/pkg/nmap"
)

// Repo defines the interface for storing and fetching Scans.
type Repo interface {
	// CreateScan inserts a Scan into the repository.
	CreateScan(context.Context, *Scan) (*Scan, error)
	// GetScan returns a single Scan by id.
	GetScan(ctx context.Context, id int64) (*Scan, error)
	// ListScans returns all Scans for the given host.
	ListScans(ctx context.Context, host string) ([]Scan, error)
	// Close closes the repository connection.
	Close() error
}

// New creates a new SQLite Repo.
func New(dbFile string) (Repo, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	r := sqliteRepo{
		dbFile: dbFile,
		DB:     db,
	}

	if err := r.createSchema(); err != nil {
		return nil, err
	}

	return &r, nil
}

// sqliteRepo implements the Repo interface against a SQLite database.
type sqliteRepo struct {
	dbFile string
	DB     *sql.DB
}

// CreateScan inserts a Scan into the scan table.
func (r *sqliteRepo) CreateScan(ctx context.Context, scan *Scan) (*Scan, error) {
	const insert = `insert into scan(host, result) values(?, ?)`

	stmt, err := r.DB.Prepare(insert)
	if err != nil {
		return nil, err
	}

	data, err := encodeScanResults(scan.Ports)
	if err != nil {
		return nil, err
	}

	resp, err := stmt.Exec(scan.Host, data)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetScan(ctx, id)
}

// GetScan returns a single Scan by ID.
func (r *sqliteRepo) GetScan(ctx context.Context, id int64) (*Scan, error) {
	const query = `SELECT host, result, created_at FROM scan WHERE id=?;`

	var (
		host   string
		result []byte
		ts     string
	)

	row := r.DB.QueryRow(query, id)

	switch err := row.Scan(&host, &result, &ts); err {
	case nil:
		// break
	case sql.ErrNoRows:
		return nil, fmt.Errorf("scan with id=%d not found", id)
	default:
		return nil, fmt.Errorf("failed to query scan: %w", err)
	}

	ports, err := decodeScanResults(result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode scan results")
	}

	createdAt, err := decodeCreatedAt(ts)
	if err != nil {
		return nil, fmt.Errorf("failed to decode created_at timestamp: %w", err)
	}

	return &Scan{
		ID:        id,
		CreatedAt: createdAt,
		ScanResult: nmap.ScanResult{
			Host:  host,
			Ports: ports,
		},
	}, nil
}

// ListScans returns all scans for a given host, ordered by most recent.
func (r *sqliteRepo) ListScans(ctx context.Context, host string) ([]Scan, error) {
	const query = `SELECT id, result, created_at FROM scan WHERE host=? ORDER BY created_at DESC;`

	row, err := r.DB.Query(query, host)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	scans := make([]Scan, 0)
	for row.Next() {
		var (
			id     int64
			result []byte
			ts     string
		)

		row.Scan(&id, &result, &ts)

		ports, err := decodeScanResults(result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode scan results")
		}

		createdAt, err := decodeCreatedAt(ts)
		if err != nil {
			return nil, fmt.Errorf("failed to decode created_at timestamp: %w", err)
		}

		scans = append(scans, Scan{
			ID: id,
			ScanResult: nmap.ScanResult{
				Host:  host,
				Ports: ports,
			},
			CreatedAt: createdAt,
		})
	}

	return scans, nil
}

// Close closes the SQLite database connection.
func (r *sqliteRepo) Close() error {
	return r.DB.Close()
}

// createSchema initializes the SQLite database schema.
func (r *sqliteRepo) createSchema() error {
	const schema = `
	CREATE TABLE IF NOT EXISTS scan(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		host TEXT NOT NULL,
		result BLOB NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_scan_host ON scan(host);`

	if _, err := r.DB.Exec(schema); err != nil {
		return err
	}

	return nil
}
