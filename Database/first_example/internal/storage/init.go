package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// OpenConnection is call one TIME!
func OpenConnection(pathDbFile string) (*sql.DB, error) {
	return sql.Open("sqlite3", pathDbFile)
}

func CreateTable(conn *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS daily_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		day TEXT NOT NULL,
		category TEXT NOT NULL,
		question TEXT NOT NULL,
		scale INTEGER NOT NULL,

		UNIQUE (day)
	);
	`
	_, err := conn.Exec(query)

	return err
}
