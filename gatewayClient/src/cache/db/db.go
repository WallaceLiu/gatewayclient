package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func MyDBOpen(dbType, dbUrl string) *sql.DB {
	db, err := sql.Open(dbType, dbUrl)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func MyDBClose(db *sql.DB) {
	if db == nil {
		log.Fatal("db is nil")
		return
	}
	db.Close()
}
