package main

import (
	"fmt"
)

func main() {
	var nums int
	var maxVal int
	var countMaxVal int
	var iteration int

	for {
		iteration++
		fmt.Scan(&nums)
		if nums == 0 {
			break
		}

		if nums > maxVal {
			maxVal = nums
			countMaxVal = 1
		} else if nums == maxVal {
			countMaxVal++
		}

		fmt.Println("iteration: #", iteration)
		fmt.Println("max val in this iteration:", maxVal)
		fmt.Println("counter max val in this iteration:", countMaxVal)
		fmt.Println()
	}
	fmt.Println("max val:", maxVal)
	fmt.Println("num max val:", countMaxVal)

}
