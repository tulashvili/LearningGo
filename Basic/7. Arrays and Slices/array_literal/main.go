package main

import "fmt"

func main() {
	var coffeeType = [3]string{"Espr", "Latte", "Cappuc"}
	fmt.Println("Types of coffee:", coffeeType)
	fmt.Println("Nubmers in massive:", len(coffeeType))

	coffeeType[len(coffeeType)-3] = "Milk"
	fmt.Println("Types of coffee:", coffeeType)

}
