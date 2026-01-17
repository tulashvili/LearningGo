package storage

import "database/sql"

func UpdateRows(conn *sql.DB) error {
	// WHERE category = "Psy" change value all items with Psy
	query := `
	UPDATE daily_log
	SET scale = 1
	WHERE id = 2
	;
	`
	_, err := conn.Exec(query)

	return err
}
