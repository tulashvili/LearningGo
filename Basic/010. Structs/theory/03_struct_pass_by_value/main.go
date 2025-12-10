package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType string
	Size       string
}

func printReceiptCopy(order CoffeeOrder) { // copy value of my order
	order.Size = "Small"
	fmt.Println("Order size in the function:", order.Size)
}

func main() {
	myOrder := CoffeeOrder{"Americano", "Large"}
	fmt.Println("My order in the main before printReceiptCopy function call:", myOrder)
	printReceiptCopy(myOrder) // doesnt affect my order
	fmt.Println("My order in the main after printReceiptCopy function call:", myOrder)
}
