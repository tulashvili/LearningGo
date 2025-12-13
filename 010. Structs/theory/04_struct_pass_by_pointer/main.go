package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType string
	Size       string
}

func printReceiptCopyPointer(order *CoffeeOrder) { // copy value of my order
	order.Size = "Small"
	fmt.Println("Order size in the function:", order.Size)

	pointerToMyOrder := &order
	fmt.Println("Pointer address in the func:", &pointerToMyOrder)
	fmt.Println("Order address in the func:", &order)
	fmt.Println("Order Size address in the func:", &order.Size)
	fmt.Println("Order CoffeeType address in the func:", &order.CoffeeType)

}

// func printReceiptCopyPassByValue(order CoffeeOrder) { // copy value of my order
// 	order.Size = "Small"
// 	fmt.Println("Order size in the function:", order.Size)
// }

func main() {
	myOrder := CoffeeOrder{"Americano", "Large"}
	fmt.Println("My order in the main before printReceiptCopy* function call:", myOrder.Size)

	// printReceiptCopyPassByValue(myOrder) // doesnt affect my order

	printReceiptCopyPointer(&myOrder)

	fmt.Println("My order in the main after printReceiptCopyPointer function call:", myOrder.Size)
	// fmt.Println("My order in the main after printReceiptCopyPassByValue function call:", myOrder.Size) // doesnt affect my order

	pointerToMyOrder := &myOrder
	fmt.Println("Pointer address in the main func:", &pointerToMyOrder)

}
