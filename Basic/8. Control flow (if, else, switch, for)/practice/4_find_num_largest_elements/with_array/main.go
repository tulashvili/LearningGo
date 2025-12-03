package main

import "fmt"

func main() {

	nums := []int{1, 3, 3, 1}
	maxVal := nums[0]
	maxNums := []int{}
	var iteration int

	for i := 0; i < len(nums); i++ {
		iteration++

		fmt.Printf("#%d: Число %d\n", iteration, nums[i])

		// Находим максимальное значение
		if nums[i] >= maxVal {
			maxVal = nums[i] // #1 - 1, #2 - 3, #3 - 3, #4 - 3
		}
		fmt.Printf("Самое большое число на текущий момент - %d\n", maxVal)

		fmt.Println()
	}
	fmt.Println("Максимальное число массива", maxVal)

	for _, v := range nums {
		if v == maxVal {
			maxNums = append(maxNums, v)
		}
	}
	fmt.Println("Количество самых больших одинаковых чисел", len(maxNums))

}
