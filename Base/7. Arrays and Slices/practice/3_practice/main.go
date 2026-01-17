package main

import "fmt"

func main() {
	var number int
	// fmt.Print("Введите количество чисел: ")
	fmt.Scan(&number)

	array := make([]int, number)

	var itemArray int
	for index := 0; index < number; index++ {
		// fmt.Print("Введите элемент массива:")
		fmt.Scan(&itemArray)
		array[index] = itemArray
		// fmt.Println(array)
	}

	// fmt.Println("Четные позиции:")
	for index := 0; index < number; index += 2 {
		if index > 0 {
			fmt.Print(" ")
		}
		fmt.Print(array[index])
	}
}
