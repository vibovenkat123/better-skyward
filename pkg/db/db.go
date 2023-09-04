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
	return db, nil
}
