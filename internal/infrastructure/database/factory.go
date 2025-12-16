package database

import (
	"errors"

	"github.com/florian-renfer/beanbase.io/internal/adapter/repository"
)

var (
	// errInvalidSQLDatabaseInstance is returned when an invalid SQL database instance is requested.
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

const (
	// InstancePostgres represents the PostgreSQL database instance type.
	InstancePostgres int = iota
)

// NewDatabaseSQLFactory returns a SQL repository based on the given instance type.
func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstancePostgres:
		return NewPostgresHandler(newConfigPostgres())
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
