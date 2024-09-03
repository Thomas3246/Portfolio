package service

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// func Check(err error)
// function checks if error is nil or not
// if it is -> fatal error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func OpenDB(path string) *sql.DB
// function gets path to the database and opens it
//
// ---ATTENTION---
// works only with sqlite3
// ---ATTENTION---
func OpenDB(path string) *sql.DB {
	database, err := sql.Open("sqlite3", path)
	Check(err)
	return (database)
}
