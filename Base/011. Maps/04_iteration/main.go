package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":  2.50,
		"Latte":     3.75,
		"Americano": 2.75,
	}

	for drink, price := range menu {
		fmt.Println(drink, price)
	}
}
