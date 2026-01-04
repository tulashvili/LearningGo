package storage

import (
	"database/sql"
	"fmt"
	"test_work_db/internal/model"
)

func SelectRow(conn *sql.DB) ([]model.Entry, error) {
	// WHERE category = "Psy" change value all items with Psy
	query := `
	SELECT * FROM daily_log 
	WHERE scale = 4
	ORDER BY id ASC
	LIMIT 100
	;
	`
	res, err := conn.Query(query)

	defer res.Close()

	currentDay := make([]model.Entry, 0)

	for res.Next() {
		var entry model.Entry

		err := res.Scan(
			&entry.ID,
			&entry.Day,
			&entry.Category,
			&entry.Question,
			&entry.Scale,
		)
		if err != nil {
			return nil, err
		}
		currentDay = append(currentDay, entry)
		formatedPrintRows(entry)
	}
	return currentDay, err
}

func formatedPrintRows(currentDay model.Entry) {
	fmt.Println("---------------------")
	fmt.Println("id:", currentDay.ID)
	fmt.Println("day:", currentDay.Day)
	fmt.Println("category:", currentDay.Category)
	fmt.Println("question:", currentDay.Question)
	fmt.Println("scale:", currentDay.Scale)
}
