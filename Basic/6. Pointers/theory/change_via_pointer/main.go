package main

import "fmt"

func main() {
	// STEP 1:
	// CompileTime (code you write): coffeePrice := 4.50
	// Runtime - what machine sees: [some memory address] = 4.50
	coffeePrice := 4.50

	// STEP 2:
	// CompileTime (code you write): fmt.Println("Coffee price:", coffeePrice)
	// Runtime - what machine sees: [some memory address] = fmt.Println([some mem address], [mem address (same as on step 1)] )
	fmt.Println("Coffee price:", coffeePrice)
	fmt.Println("Mem address for 4.50:", &coffeePrice)

	coffeePrice = 5
	fmt.Println("Mem address for 5:", &coffeePrice)

	var ponterToCoffeePrice *float64 = &coffeePrice
	// go to the memory location where ponterToCoffeePrice
	// and change value in this memory location
	*ponterToCoffeePrice = 5.50
	fmt.Println("Mem address for 5.50 (change with pointer):", coffeePrice)
}
