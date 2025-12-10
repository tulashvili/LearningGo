package main

import "fmt"

func main() {
	var length int
	fmt.Println("Введите количество элементов в слайсе")
	fmt.Scan(&length)

	var getElemFromSlice int
	fmt.Println("Введите, какой элемент из слайса хочешь получить элементов в слайсе:")
	fmt.Scan(&getElemFromSlice)

	mySlice := make([]int, length)
	for i := range mySlice {
		fmt.Scan(&mySlice[i])
	}

	fmt.Print(mySlice[getElemFromSlice])
}
