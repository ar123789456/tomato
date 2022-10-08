package server

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	log.Println("Init DB")
	db, err := sql.Open("sqlite3", "test")
	if err != nil {
		log.Fatalln(err)
	}
	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("SQLITE_VERSION:", version)
	return db
}
