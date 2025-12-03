package main

import "fmt"

func getSmallestNum(nums [4]int) int {
	minVal := nums[0]
	var iteration int

	for _, v := range nums {
		iteration++
		fmt.Println("iteration: #", iteration)

		if minVal > v {
			minVal = v
		}

		fmt.Println("num in this iteration", v)
		fmt.Println("smallest minValue in this iteration", minVal)
		fmt.Println()
	}
	return minVal
}

func main() {
	nums := [4]int{4, 5, 6, 7}
	fmt.Println("smallest minValue", getSmallestNum(nums))
}
