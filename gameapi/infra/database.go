package infra

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/task4233/techtrain-mission/gameapi/config"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// NewDB connects database and pointer for sqlx.DB
// If opening DB with this function, Do Close() like example
func NewDB() (*sqlx.DB, error) {
	// this process might be moved config package
	var DSN string = config.DSN()

	db, err := sqlx.Open("mysql", DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	var version string
	if err = db.Get(&version, "SELECT version()"); err != nil {
		return nil, fmt.Errorf("failed to Get(accessed DSN=> \"%s\"): %w", DSN, err)
	}
	if version == "" {
		return nil, fmt.Errorf("failed to get MySQL version")
	}

	return db, nil
}
