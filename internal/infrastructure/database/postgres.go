package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/florian-renfer/beanbase.io/internal/adapter/repository"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// postgresHandler handles PostgreSQL database operations.
type postgresHandler struct {
	db *sql.DB
}

// NewPostgresHandler returns a SQL repository for a PostgreSQL instance.
func NewPostgresHandler(c *config) (*postgresHandler, error) {
	var dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	db, err := sql.Open(c.driver, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	performMigration(c)

	return &postgresHandler{db: db}, nil
}

// performMigration applies database schema migrations using the provided configuration.
// It ensures that the database is up to date with the latest schema changes.
func performMigration(c *config) {
	migrationDsn := fmt.Sprintf("pgx5://%s:%s@%s:%s/%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	m, err := migrate.New("file://db/migrations/postgres", migrationDsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[migrate] Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err := m.Up(); err != nil {
		fmt.Fprintf(os.Stderr, "[migrate] Unable to perform database migration: %v\n", err)
		os.Exit(1)
	}
}

// BeginTx starts a new database transaction.
func (p postgresHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return postgresTx{}, err
	}

	return newPostgresTx(tx), nil
}

// ExecuteContext executes a query without returning any rows.
func (p postgresHandler) ExecuteContext(ctx context.Context, query string, args ...any) error {
	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// QueryContext executes a query that returns rows.
func (p postgresHandler) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newPostgresRows(rows)

	return row, nil
}

// QueryRowContext executes a query that is expected to return at most one row.
func (p postgresHandler) QueryRowContext(ctx context.Context, query string, args ...any) repository.Row {
	row := p.db.QueryRowContext(ctx, query, args...)

	return newPostgresRow(row)
}

// postgresRow wraps a single SQL row.
type postgresRow struct {
	row *sql.Row
}

// newPostgresRow creates a new postgresRow.
func newPostgresRow(row *sql.Row) postgresRow {
	return postgresRow{row: row}
}

// Scan copies the columns from the row into the destination values.
func (pr postgresRow) Scan(dest ...any) error {
	if err := pr.row.Scan(dest...); err != nil {
		return err
	}

	return nil
}

// postgresRows wraps multiple SQL rows.
type postgresRows struct {
	rows *sql.Rows
}

// newPostgresRows creates a new postgresRows.
func newPostgresRows(rows *sql.Rows) postgresRows {
	return postgresRows{rows: rows}
}

// Scan copies the columns from the current row into the destination values.
func (pr postgresRows) Scan(dest ...any) error {
	if err := pr.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

// Next prepares the next result row for reading.
func (pr postgresRows) Next() bool {
	return pr.rows.Next()
}

// Err returns the error, if any, that was encountered during iteration.
func (pr postgresRows) Err() error {
	return pr.rows.Err()
}

// Close closes the rows iterator.
func (pr postgresRows) Close() error {
	return pr.rows.Close()
}

// postgresTx wraps a SQL transaction.
type postgresTx struct {
	tx *sql.Tx
}

// newPostgresTx creates a new postgresTx.
func newPostgresTx(tx *sql.Tx) postgresTx {
	return postgresTx{tx: tx}
}

// ExecuteContext executes a query within the transaction.
func (p postgresTx) ExecuteContext(ctx context.Context, query string, args ...any) error {
	_, err := p.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// QueryContext executes a query within the transaction that returns rows.
func (p postgresTx) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	rows, err := p.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newPostgresRows(rows)

	return row, nil
}

// QueryRowContext executes a query within the transaction that is expected to return at most one row.
func (p postgresTx) QueryRowContext(ctx context.Context, query string, args ...any) repository.Row {
	row := p.tx.QueryRowContext(ctx, query, args...)

	return newPostgresRow(row)
}

// Commit commits the transaction.
func (p postgresTx) Commit() error {
	return p.tx.Commit()
}

// Rollback aborts the transaction.
func (p postgresTx) Rollback() error {
	return p.tx.Rollback()
}
