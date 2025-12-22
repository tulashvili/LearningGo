package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// вводим тип идентично полям, которые у нас есть в users.json
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", data) // []uint8 → []byte

	var user User
	// json принимает только байты!!!
	err = json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Имя пользователя: %s,\nВозраст пользователя: %d\n", user.Name, user.Age)
}
