package main

import "fmt"

func main() {
	// slice because [] - is empty
	ratings := []int{5, 4, 3, 2}
	fmt.Println("original rate:", ratings)

	ratings[2] = 1
	fmt.Println("Change rate with rating = 3:", ratings)

	coffeeType := make([]string, 3)
	fmt.Println("Coffee type:", coffeeType)
	fmt.Println("Len coffee type:", len(coffeeType))

	coffeeType[1] = "Banana Raf"
	coffeeType[2] = "Americano"
	fmt.Println("Coffee type after reassignment:", coffeeType)

}
