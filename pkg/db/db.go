package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func Open() (*sql.DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/.better-skyward/db.db", home))
	if err != nil {
		return nil, err
	}
	// set a schema
	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS classes (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL,
    	teacher TEXT NOT NULL,
    	room TEXT NOT NULL,
    	period INTEGER NOT NULL
);`)

	if err != nil {
		return nil, err
	}
	return db, nil
}
