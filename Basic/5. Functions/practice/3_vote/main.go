package main

import "fmt"

func vote(x int, y int, z int) int {
	sumInt := x + y + z
	// var oftenInt int

	for i := 0; i <= sumInt; i++ {
		if sumInt >= 2 {
			return 1
		} else {
			return 0
		}
	}
	return sumInt
}
func main() {
	fmt.Println(vote(0, 0, 1))
}

// func main() {
// 	vote(0, 1, 1)
// }
