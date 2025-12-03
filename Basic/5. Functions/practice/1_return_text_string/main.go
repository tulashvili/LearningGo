package main

import "fmt"

func f(text string) string {
	fmt.Println(text)
	return text
}
func main() {
	f("Hello!")
}
