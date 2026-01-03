package storage

import (
	"database/sql"
	"fmt"
)

func SelectRow(conn *sql.DB) error {
	// WHERE category = "Psy" change value all items with Psy
	query := `
	SELECT * FROM daily_log 
	WHERE scale = 4
	ORDER BY id ASC
	LIMIT 100
	;
	`
	res, err := conn.Query(query)

	// defer res.Close()

	for res.Next() {
		var id int
		var day string // because in sqlite it's TEXT type
		var category string
		var question string
		var scale int

		err := res.Scan(
			&id,
			&day,
			&category,
			&question,
			&scale,
		)
		if err != nil {
			return err
		}

		formatedPrintRows(id, day, category, question, scale)
	}
	return err
}

func formatedPrintRows(
	id int,
	day string,
	category string,
	question string,
	scale int,
) {
	fmt.Println("---------------------")
	fmt.Println("id:", id)
	fmt.Println("day:", day)
	fmt.Println("category:", category)
	fmt.Println("question:", question)
	fmt.Println("scale:", scale)
}
