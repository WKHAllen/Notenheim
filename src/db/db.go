package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// fixQueryParams turns a query containing '?' characters into '$1, $2, $3, etc.'
func fixQueryParams(sql string) string {
	for paramCount := 1; strings.Contains(sql, "?"); paramCount++ {
		sql = strings.Replace(sql, "?", fmt.Sprintf("$%v", paramCount), 1)
	}
	return sql
}

// Manager represents a database manager
type Manager struct {
	dbURL string
	pool  *pgxpool.Pool
}

// NewDBManager creates a new database manager object
func NewDBManager(dbURL string) (*Manager, error) {
	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return &Manager{
		dbURL: dbURL,
		pool:  dbpool,
	}, nil
}

// Execute executes a sql query
func (dbm *Manager) Execute(sql string, params ...interface{}) error {
	_, err := dbm.pool.Exec(context.Background(), fixQueryParams(sql), params...)
	return err
}

// QueryRow executes a sql query that returns a row
func (dbm *Manager) QueryRow(sql string, params ...interface{}) pgx.Row {
	return dbm.pool.QueryRow(context.Background(), fixQueryParams(sql), params...)
}

// QueryRows executes a sql query that returns multiple rows
func (dbm *Manager) QueryRows(sql string, params ...interface{}) (pgx.Rows, error) {
	return dbm.pool.Query(context.Background(), fixQueryParams(sql), params...)
}

// Close closes the database pool
func (dbm *Manager) Close() {
	dbm.pool.Close()
}
