package main

import (
	"encoding/json"
	"os"
)

// вводим тип идентично полям, которые у нас есть в users.json
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "Omar",
		Age:  27,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("user_out.json", data, 0644)
	if err != nil {
		panic(err)
	}

	user.Name = "Jack"
	data, err = json.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("user_out.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
