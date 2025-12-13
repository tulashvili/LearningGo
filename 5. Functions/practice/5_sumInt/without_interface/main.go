package main

import "fmt"

func sumInt(nums ...int) (a int, b int) {

	var sum int
	var count int

	for _, elem := range nums {
		sum += elem
		count++
	}
	return count, sum
}

func main() {
	a, b := sumInt(1, 0)
	fmt.Printf("%d, %d\n", a, b)

}
