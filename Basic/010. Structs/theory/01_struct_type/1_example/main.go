package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType       string
	CoffeeSize       string
	CustomerName     string
	CustomerYear     int
	CustomerIsMember bool
}

func main() {
	var order CoffeeOrder
	fmt.Println(order)

	order.CoffeeType = "Americano"
	order.CoffeeSize = "Medium"
	order.CustomerName = "Omar"
	order.CustomerYear = 1998
	order.CustomerIsMember = true

	fmt.Println(order)
	fmt.Println(order.CoffeeType)

}
