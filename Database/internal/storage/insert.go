package storage

import (
	"database/sql"
	"time"
)

func InsertRows(
	conn *sql.DB,
	day time.Time,
	category string,
	question string,
	scale int,
) error {
	query := `
	INSERT INTO daily_log (day, category, question, scale)
	VALUES ($1, $2, $3, $4)
	;
	`
	_, err := conn.Exec(query, day, category, question, scale)

	return err
}
