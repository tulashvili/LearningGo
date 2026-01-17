package main

import "fmt"

func main() {
	tokens := 3

	for tokens > 0 {
		// infinite cicle without tokens
		tokens--
		fmt.Println("Make another coffee cups ....")
	}
}
