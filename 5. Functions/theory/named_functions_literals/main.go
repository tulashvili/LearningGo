package main

import "fmt"

func main() {
	taxRate := 0.10 // 10% tax

	calculateTax := func(amount float64) float64 {
		return amount * taxRate
	}

	subTotal := 25.00
	tax := calculateTax(subTotal)
	total := tax + subTotal
	fmt.Printf("Total amount = $%.2f\n", total)
}
