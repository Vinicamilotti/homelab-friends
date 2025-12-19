package store

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/duckdb/duckdb-go/v2"
)

type SQLDBConn struct {
	Path string
}

type DBConnectionFacade interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Begin() (*sql.Tx, error)
	Close() error
}

type SLQDBError struct {
	Error error
}

func (e *SLQDBError) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, e.Error
}

func (e *SLQDBError) Exec(query string, args ...any) (sql.Result, error) {
	return nil, e.Error
}

func (e *SLQDBError) Begin() (*sql.Tx, error) {
	return nil, e.Error
}

func (e *SLQDBError) Close() error {
	return e.Error
}
func (e *SLQDBError) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return nil, e.Error
}
func (e *SLQDBError) Prepare(query string) (*sql.Stmt, error) {
	return nil, e.Error
}

func (e *SLQDBError) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, e.Error
}

func NewSqliteConn() SQLDBConn {
	path := os.Getenv("DB_PATH")
	return SQLDBConn{Path: path}
}
func (s *SQLDBConn) DB() DBConnectionFacade {
	db, err := sql.Open("duckdb", s.Path)
	if err != nil {
		log.Fatal("CANNOT CONNECT TO SQLITE: ", err)
		return &SLQDBError{Error: err}
	}
	return db
}
