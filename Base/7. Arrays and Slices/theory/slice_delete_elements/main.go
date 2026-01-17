package main

import (
	"fmt"
)

func deleteItemFromMenu(index int, slice []string) []string {
	// return slices.Delete(slice, index, index+1)
	return append(slice[:index], slice[index+1:]...)

}

func main() {
	coffeeType := []string{"Espresso", "Flat", "Americano", "Mocha"}
	fmt.Println("Original menu:", coffeeType)
	fmt.Println("Len slice:", len(coffeeType), "capacity slice:", cap(coffeeType))

	fmt.Println()

	indexToRemove := 1
	coffeeType = deleteItemFromMenu(indexToRemove, coffeeType)
	fmt.Println("New menu:", coffeeType)
	fmt.Println("Len slice:", len(coffeeType), "capacity slice:", cap(coffeeType))

}
