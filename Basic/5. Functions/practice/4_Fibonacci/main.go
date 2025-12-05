package main

import "fmt"

func main() {
	var iteration int
	var previousVal_1 = 1
	var previousVal_2 = 1
	var nextValue int

	for i := 0; i < 4; i++ {
		iteration++
		fmt.Println("iteration #", iteration)

		nextValue = previousVal_1 + previousVal_2
		previousVal_1 = previousVal_2
		previousVal_2 = nextValue

		fmt.Println("next value:", nextValue)

		fmt.Println()
	}
}

// func fibonacci(n int) int {
// 	n = 1
// 	var g int
// 	for i := 0; i < 10; i++ {
// 		g += n
// 	}
// 	return g
// }
