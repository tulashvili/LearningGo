package main

import (
	"log"
	"test_work_db/internal/storage"
)

func main() {
	pathDbFile := "test.db"

	db, err := storage.OpenConnection(pathDbFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.CreateTable(db); err != nil {
		log.Fatal(err)
	}

	// if err := storage.InsertRows(
	// 	db,
	// 	time.Now(),
	// 	"Psy",
	// 	"Как твое психологическое состояние?",
	// 	4,
	// ); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Данные добавлены успешно в таблицу daily_log")

	// if err := storage.UpdateRows(db); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Данные обновлены в таблице daily_log")

	if err := storage.DeleteRow(db); err != nil {
		log.Fatal(err)
	}

}
