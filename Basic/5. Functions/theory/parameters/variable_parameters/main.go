package main

import "fmt"

func myPrint(a ...interface{}) {
	for _, elem := range a {
		fmt.Printf("%v ", elem)
	}
}

func main() {
	myPrint(1, true, "hi", 123, 3.41)
}
