package main

import (
	"fmt"
	"log"
	"test_work_db/internal/storage"
)

func main() {
	pathDbFile := "test.db"

	db, err := storage.OpenConnection(pathDbFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Подключение к базе установлено успешно")

	if err := storage.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Таблица в базе данных была создана успешно")

	if err := storage.InsertRows(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Данные добавлены успешно в таблицу daily_log")

}
