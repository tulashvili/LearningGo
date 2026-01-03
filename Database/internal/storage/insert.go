package storage

import (
	"database/sql"
)

func InsertRows(conn *sql.DB) error {
	query := `
	INSERT INTO daily_log (day, category, question, scale)
	VALUES ('2025-11-12','Bio','Покушал сегодня?',4)
	;
	`
	_, err := conn.Exec(query)

	return err
}
