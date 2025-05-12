package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("sqlite", "./bookmarks.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database")
	migrate()
}

func CloseDB() {
	if err := Db.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("database closed")
}
