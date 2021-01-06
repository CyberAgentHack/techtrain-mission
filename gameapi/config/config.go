package config

import (
	"fmt"
	"os"
)

// Port returns used port in .env
func Port() string {
	return os.Getenv("PORT")
}

// DSN returns DataSourceName in .env
func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	) + "?parseTime=true&collation=utf8mb4_bin"
}

// IsDev returns for development or not
func IsDev() bool {
	return os.Getenv("ENV") == "dev"
}

// IsTest returns for test or not
func IsTest() bool {
	return os.Getenv("ENV") == "test"
}
