package main

import "fmt"

func main() {
	var number int
	// fmt.Print("Введите количество чисел в массиве: ")
	fmt.Scan(&number)

	array := make([]int, number)

	for index := 0; index < number; index++ {
		fmt.Scan(&array[index])
	}

	countPositive := 0
	for index := 0; index < len(array); index++ {
		if array[index] > 0 {
			countPositive++
		}
	}
	fmt.Println("Количество положительных чисел в массиве:", countPositive)
}
