package storage

import (
	"database/sql"
	"fmt"
)

func DeleteRow(conn *sql.DB) error {
	// WHERE category = "Psy" change value all items with Psy
	query := `
	DELETE from daily_log
	WHERE id = 4
	;
	`
	res, err := conn.Exec(query)

	affected, err := res.RowsAffected()

	if affected == 0 {
		fmt.Println("Запись не найдена")
	} else {
		fmt.Println("Запись удалена успешно")
	}

	return err
}
