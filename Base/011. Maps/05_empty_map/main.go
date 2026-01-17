package main

import "fmt"

func main() {
	var stock map[string]int
	fmt.Println("Inside map:", stock)
	fmt.Println("Location in memory", &stock)
	fmt.Printf("Location in memory %p\n", &stock) // var address, map is nil

	if stock == nil {
		fmt.Println("Map is nil")
	}

	stock["Latte"] = 10 // panic: assignment to entry in nil map
	fmt.Println(stock)
}
