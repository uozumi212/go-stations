package db

import (
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

// NewDB returns go-sqlite3 driver based *sql.DB.
func NewDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path+"?parseTime=true&loc=Local")
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
