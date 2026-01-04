package storage

import (
	"database/sql"
	"test_work_db/internal/model"
)

func InsertRows(
	conn *sql.DB,
	e model.Entry,
) error {
	query := `
	INSERT INTO daily_log (day, category, question, scale)
	VALUES ($1, $2, $3, $4)
	;
	`
	_, err := conn.Exec(query, e.Day, e.Category, e.Question, e.Scale)

	return err
}
