package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db       *sql.DB
	Products ProductsTable
}

func NewDatabase(file string) (empty Database, err error) {
	var db *sql.DB
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		return empty, err
	}

	return Database{
		db:       db,
		Products: NewProductsTable(db),
	}, nil
}
