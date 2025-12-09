package main

import "fmt"

func test(x1 *int, x2 *int) {
	*x1 = *x1 * *x2
	fmt.Println(*x1)
}

func main() {
	x1 := 2
	x2 := 2
	test(&x1, &x2)
}
