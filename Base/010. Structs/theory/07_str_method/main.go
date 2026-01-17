package main

import "fmt"

type CoffeeType string

func (coffeeType CoffeeType) Describe() {
	fmt.Printf("%s - This is delicious classical coffee type\n", coffeeType)
}
func main() {
	var myCoffee CoffeeType = "Americano"
	myCoffee.Describe()

}
