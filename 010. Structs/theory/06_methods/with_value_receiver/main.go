package main

import "fmt"

// Так делать нужно!
type CoffeeShop struct {
	Name string
}

// method with value receiver
func (shop CoffeeShop) GreetShop() {
	fmt.Println("Welcome to", shop.Name)
}

func main() {
	myShop := CoffeeShop{
		Name: "From SoSo Coffee with Grains",
	}
	myShop.GreetShop()
}
