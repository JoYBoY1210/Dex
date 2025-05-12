package db

import (
	"log"
	// "github.com/mattn/go-sqlite3"
	// "github.com/JoYBoY12/Dex/backend/database"
)

func migrate() {

	_, err := Db.Exec(`CREATE TABLE IF NOT EXISTS bookmarks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		url TEXT NOT NULL,
		pinned BOOLEAN NOT NULL DEFAULT 0
	)`)

	if err != nil {
		log.Fatalf("failed to create bookmarks table: %v", err)
	}
	// _, err = Db.Exec(`ALTER TABLE bookmarks ADD COLUMN category TEXT DEFAULT ''`)


	log.Println("migrated database")
}
