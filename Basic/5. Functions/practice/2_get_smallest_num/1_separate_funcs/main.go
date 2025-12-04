package main

import "fmt"

func getSmallestNum(nums []int) int {
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
	// declare len slice
	var number = 4
	array := make([]int, number)

	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(getSmallestNum([]int(array)))
}
