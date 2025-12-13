package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType string
	Size       string
}

func main() {
	order := CoffeeOrder{
		CoffeeType: "Americano",
		Size:       "Large",
	}
	fmt.Println(order)
	fmt.Println(order.CoffeeType)
	fmt.Println(order.Size)
	fmt.Println()

	myFriendsOrder := CoffeeOrder{"Espresso", "Small"}

	fmt.Println(myFriendsOrder)
	fmt.Println(myFriendsOrder.CoffeeType)
	fmt.Println(myFriendsOrder.Size)
}
