package main

import "fmt"

// Так делать не нужно!
type CoffeeShop struct {
	Name  string
	Greet func(shop CoffeeShop)
}

func anotherGreet(shop CoffeeShop) {
	fmt.Println("Hello", shop.Name)
}

func greetShop(shop CoffeeShop) {
	fmt.Println("Welcome to", shop.Name)
}

func main() {
	myShop := CoffeeShop{
		Name: "From SoSo Coffee with Grains",
		// Greet: anotherGreet,
		Greet: greetShop,
	}
	myShop.Greet(myShop)
}
