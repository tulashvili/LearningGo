package main

import "fmt"

func main() {
	coffeeCups := 10

	// Preparing coffee cup #4
	// Preparing coffee cup #5
	// Preparing coffee cup #6
	// Preparing coffee cup #7
	// Preparing coffee cup #8
	// Preparing coffee cup #9
	// for i := 4; i < coffeeCups; i++ {
	// 	fmt.Printf("Preparing coffee cup #%d\n", i)
	// }

	// Preparing coffee cup #2
	// Preparing coffee cup #3
	// Preparing coffee cup #4
	// Preparing coffee cup #5
	// Preparing coffee cup #6
	// Preparing coffee cup #7
	// Preparing coffee cup #8
	// Preparing coffee cup #9
	// for i := 2; i < coffeeCups; i++ {
	// 	fmt.Printf("Preparing coffee cup #%d\n", i)

	// }
	// 1,2...10
	for i := 1; i <= coffeeCups; i++ {
		fmt.Printf("Preparing coffee cup #%d\n", i)
	}

	fmt.Println()

	// 10,9 .... 1
	for i := coffeeCups; i >= 0; i-- {
		fmt.Printf("Preparing coffee cup #%d\n", i)
	}
}
