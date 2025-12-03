package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Пюрешка", "Котлетка", "Какашка", "Водичка"}
	fmt.Println("Dessert Menu:", dessertMenu)

	slice := dessertMenu[:] // all elem
	fmt.Println("Slice from dessert menu", slice)

	slice = dessertMenu[2:]
	fmt.Println("All elem from 2", slice)

	slice = dessertMenu[:3]
	fmt.Println("All elem before 3", slice)
}
