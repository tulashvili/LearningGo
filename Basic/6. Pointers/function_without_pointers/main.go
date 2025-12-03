package main

import "fmt"

func applyDiscount(price float64, discountRate float64) float64 {
	return price - (price * discountRate)
}
func main() {
	var coffeePrice float64 = 5.00
	discount := 0.10
	fmt.Printf("Basic Coffee Price:$%.2f\n", coffeePrice)

	coffeePrice = applyDiscount(coffeePrice, discount)
	fmt.Printf("Price after discount:$%.2f\n", coffeePrice)
}
