package main

import "fmt"

func main() {
	a := 200
	b := &a // pointer to a
	*b++    // var b *int = &a
	fmt.Println(a)
}
