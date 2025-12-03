package main

import "fmt"

func main() {
	coffee := "Espresso"
	pointer := &coffee
	fmt.Println("Coffee name:", coffee)
	fmt.Println("Memory location:", &coffee)
	fmt.Printf("Pointer address: %p\n", pointer)
	fmt.Println("-------another memory block----------")

	// another memory block
	coffeeCopy := coffee
	fmt.Println("Coffee name:", coffeeCopy)
	fmt.Println("Memory location:", &coffeeCopy)
	fmt.Printf("Pointer address: %p\n", &coffeeCopy)

}
