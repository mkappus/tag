package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Store struct {
		*sql.DB
	}
)

var db *Store

func New(path string) (*Store, error) {
	var err error
	if db != nil {
		return db, nil
	}

	d, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db = &Store{d}
	return db, nil
}

func (db *Store) Stmt(query string) func(...interface{}) (sql.Result, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil
	}
	return func(args ...interface{}) (sql.Result, error) {
		defer tx.Commit()
		stmt, err := tx.Prepare(query)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		return stmt.Exec(args...)
	}
}

func (db *Store) StmtExec(query string, args ...interface{}) (sql.Result, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	defer tx.Commit()
	return stmt.Exec(args...)
}
