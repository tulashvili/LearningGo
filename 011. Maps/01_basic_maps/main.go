package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":  2.50,
		"Latte":     3.75,
		"Americano": 2.75,
	}
	fmt.Println(menu)
	fmt.Println("Americano costs:", menu["Americano"])

	menu["Americano"] = 3
	fmt.Println("New price Americano:", menu["Americano"])

	fmt.Println("Menu quantity:", len(menu))
	
	menu["Water"] = 0.00
	fmt.Println("Menu after added new item:", menu)
	fmt.Println("New Menu quantity:", len(menu))

}
