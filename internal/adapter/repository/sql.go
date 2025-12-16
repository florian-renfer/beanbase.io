package repository

import "context"

// SQL defines the interface for SQL database operations.
type SQL interface {
	ExecuteContext(context.Context, string, ...any) error
	QueryContext(context.Context, string, ...any) (Rows, error)
	QueryRowContext(context.Context, string, ...any) Row
	BeginTx(ctx context.Context) (Tx, error)
}

// Rows represents a set of database rows returned by a query.
type Rows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Close() error
}

// Row represents a single database row.
type Row interface {
	Scan(dest ...any) error
}

// Tx defines the interface for a database transaction.
type Tx interface {
	ExecuteContext(context.Context, string, ...any) error
	QueryContext(context.Context, string, ...any) (Rows, error)
	QueryRowContext(context.Context, string, ...any) Row
	Commit() error
	Rollback() error
}
