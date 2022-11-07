package server

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	createUserTable(db)
	createTomatoTable(db)
	createTagsTable(db)
	createTimerTable(db)
	createTagsTomatoesTable(db)
	log.Println("SQLITE_VERSION:", version)
	return db
}

func createUserTable(db *sql.DB) {
	usersTable := `
		CREATE TABLE if not exists users (
        id TEXT NOT NULL PRIMARY KEY,
        "name" TEXT,
        "secondName" TEXT,
        "nick" TEXT,
        "email" TEXT,
        "photo" TEXT,
        "class" TEXT,
        "session" TEXT
        );`
	query, err := db.Prepare(usersTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User table created successfully!")
}

func createTomatoTable(db *sql.DB) {
	tomatoTable := `
		CREATE TABLE if not exists tomatoes (
        id TEXT NOT NULL PRIMARY KEY,
        timeStart INTEGER,
        createTime INTEGER,
        title TEXT,
        context TEXT,
        user_id TEXT,
        foreign key (user_id) references users(id) on delete cascade 
	);`
	query, err := db.Prepare(tomatoTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tomatoes table created successfully!")
}

func createTagsTable(db *sql.DB) {
	tagsTable := `
		CREATE TABLE if not exists tags (
        id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        tag TEXT
	);`
	query, err := db.Prepare(tagsTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tags table created successfully!")
}

func createTagsTomatoesTable(db *sql.DB) {
	tagsTable := `
		CREATE TABLE if not exists tag_tomato (
        id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        tomatoId TEXT,
        tagId integer,
        foreign key (tagId) references  tags(id) on delete cascade, 
        foreign key (tomatoId) references tomatoes(id) on delete cascade
	);`
	query, err := db.Prepare(tagsTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tag_tomatoes table created successfully!")
}

func createTimerTable(db *sql.DB) {
	timerTable := `
		CREATE TABLE if not exists timer (
        id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        workTime integer,
        restTime integer,
        tomatoId text,
        foreign key (tomatoId) references tomatoes(id) on delete cascade
	);`
	query, err := db.Prepare(timerTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("timer table created successfully!")
}
