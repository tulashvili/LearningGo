package main

import (
	"fmt"
	"os"
)

func main() {
	dataFor := []byte("236")
	os.WriteFile("test.txt", dataFor, 0600)
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
