package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int) // create pointer
	one(xPtr)
	fmt.Println(xPtr) // address cell memory
	fmt.Println(*xPtr)
}
