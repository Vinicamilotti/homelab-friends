package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/duckdb/duckdb-go/v2"
)

type SqliteConn struct {
	Path string
}

type SqliteDBFacade interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
	Begin() (*sql.Tx, error)
	Close() error
}

type SqliteDBError struct {
	Error error
}

func (e *SqliteDBError) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, e.Error
}

func (e *SqliteDBError) Exec(query string, args ...any) (sql.Result, error) {
	return nil, e.Error
}

func (e *SqliteDBError) Begin() (*sql.Tx, error) {
	return nil, e.Error
}

func (e *SqliteDBError) Close() error {
	return e.Error
}

func NewSqliteConn() SqliteConn {
	path := os.Getenv("DUCKDB_PATH")
	return SqliteConn{Path: path}
}
func (s *SqliteConn) DB() SqliteDBFacade {
	db, err := sql.Open("duckdb", s.Path)
	if err != nil {
		log.Fatal("CANNOT CONNECT TO SQLITE: ", err)
		return &SqliteDBError{Error: err}
	}
	return db
}
