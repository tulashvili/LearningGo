package main

import (
	"fmt"
)

func main() {
	workArray := [10]int{99, 151, 137, 71, 117, 187, 20, 93, 187, 67}
	index := [6]int{1, 2, 3, 5, 7, 8}

	slice := workArray[:] // add all elements to slice
	slice = append(slice, index[:]...)

	// change places
	slice[index[0]], slice[index[1]] = slice[index[1]], slice[index[0]]
	slice[index[2]], slice[index[3]] = slice[index[3]], slice[index[2]]
	slice[index[4]], slice[index[5]] = slice[index[5]], slice[index[4]]
	// fmt.Println(slice)

	slice = slice[:10]

	for index, item := range slice {
		if index > 0 {
			fmt.Print(" ") // Печатаем пробел перед каждым элементом, начиная со второго
		}
		fmt.Print(item)
	}

}
