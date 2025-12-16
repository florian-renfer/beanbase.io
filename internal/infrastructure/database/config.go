package database

import (
	"os"
)

// config holds the database configuration values.
type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string
}

// newConfigPostgres creates a new config instance using environment variables for PostgreSQL.
func newConfigPostgres() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
