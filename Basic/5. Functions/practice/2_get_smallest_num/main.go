package main

import "fmt"

//	func getSmallestNum([]int) int {
//		nums := []int{}
//		for i := 0; i < len(nums); i++ {
//			fmt.Println(i)
//		}
//
// return i
// }
func main() {
	nums := []int{4, 2, 1, 7}

	val := nums[0] // 0
	var iteration int
	for _, v := range nums {
		iteration++
		fmt.Println("iteration: #", iteration)
		fmt.Println("val in this iteration", val)
		if val > v {
			val = v
		}

		fmt.Println("num in this iteration", v)
		fmt.Println("smallest value", val)
		fmt.Println()
	}

}
