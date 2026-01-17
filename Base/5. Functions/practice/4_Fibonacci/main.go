package main

import "fmt"

func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		previousVal1 := 1
		previousVal2 := 1
		var nextValue int

		for i := 3; i <= n; i++ {
			nextValue = previousVal1 + previousVal2
			previousVal1 = previousVal2
			previousVal2 = nextValue
		}

		return nextValue
	}
}

func main() {
	fmt.Println(getFibonacci(4))
}
