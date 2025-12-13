package main

import "fmt"

// попробовать запустить так, если что - поменять package на main
func main() {
	var order string = "Latte"
	fmt.Println("Order is:", order)

	order = "Frapuccino"
	fmt.Println("Updated order is:", order)
}
