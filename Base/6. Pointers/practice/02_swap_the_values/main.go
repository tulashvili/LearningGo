package main

import "fmt"

func test(pointerToX1 *int, pointerToX2 *int) {
	fmt.Println("Values:", *pointerToX1, *pointerToX2)
	fmt.Println("Memory cell:", &pointerToX1, &pointerToX2)
	fmt.Println()

	*pointerToX1, *pointerToX2 = *pointerToX2, *pointerToX1
	fmt.Println("Values:", *pointerToX1, *pointerToX2)
	fmt.Println("Memory cell:", pointerToX1, pointerToX2)
}

func main() {
	x1 := 2
	x2 := 4
	test(&x1, &x2)
}
