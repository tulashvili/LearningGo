package main

import "fmt"

// Так делать не нужно!
type CoffeeShop struct {
	Name  string
	Greet func(shop CoffeeShop)
}

func greetShop(shop CoffeeShop) {
	fmt.Println("Welcome to", shop.Name)
}

func main() {
	myShop := CoffeeShop{
		Name:  "From SoSo Coffee with Grains",
		Greet: greetShop,
	}
	myShop.Greet(myShop)
}
