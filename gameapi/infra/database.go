package infra

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// NewDB connects database and pointer for sqlx.DB
// If opening DB with this function, Do Close() like example
func NewDB() (*sqlx.DB, error) {
	// this process might be moved config package
	var DSN string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	) + "?parseTime=true&collation=utf8mb4_bin"

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
