package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":  2.50,
		"Latte":     3.75,
		"Americano": 2.75,
	}
	fmt.Println("Current menu:", menu)

	delete(menu, "Americano")
	fmt.Println("Updated menu:", menu)

	delete(menu, "Latte")
	fmt.Println("Updated menu:", menu)

	delete(menu, "Test")
	fmt.Println("Updated menu:", menu) // if key is not exist, map print without error
}
