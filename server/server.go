package server

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	log.Println("Init DB")
	db, err := sql.Open("sqlite3", "data/test.db")
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

func Run() {
	db := InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	mux := http.NewServeMux()
	err := http.ListenAndServe("localhost:8080", mux)
	log.Println(err)
}
