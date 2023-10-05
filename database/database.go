package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb(url string) error {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS foo (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            bar TEXT NOT NULL,
        )`)
	if err != nil {
		return err
	}

	Db = db

	return nil
}
