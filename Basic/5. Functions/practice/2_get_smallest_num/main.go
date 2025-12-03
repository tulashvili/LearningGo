package main

import "fmt"

func getSmallestNum(nums [4]int) int {
	minVal := nums[0]
	var iteration int

	for _, v := range nums {
		iteration++

		if minVal > v {
			minVal = v
		}
	}
	return minVal
}

func main() {
	nums := [4]int{4, 5, 6, 7}
	fmt.Println(getSmallestNum(nums))
}
