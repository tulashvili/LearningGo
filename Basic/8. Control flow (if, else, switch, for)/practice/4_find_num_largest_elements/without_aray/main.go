package main

import "fmt"

func main() {
	var n int
	var maxVal int
	var count int

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		if n > maxVal {
			maxVal = n
			count = 1
		} else if n == maxVal {
			count++
		}
	}

	fmt.Println(count)
}
