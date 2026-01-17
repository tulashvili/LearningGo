package main

import "fmt"

// вместо этого мы можем использовать встроенный interface String
// type Ssstring interface {
// 	Sstring()
// }

type Order struct {
	Customer string
	Item     string
	Quantity int
}

// используем встроенный interface String
func (o Order) String() string {
	return fmt.Sprintf("Order: %s has ordered %s (%d)", o.Customer, o.Item, o.Quantity)
}
func main() {
	order := Order{
		Customer: "Omar",
		Item:     "Latte",
		Quantity: 2,
	}

	// используем встроенный interface String
	// fmt.Println(order.Sstring())
	fmt.Println(order)
}
