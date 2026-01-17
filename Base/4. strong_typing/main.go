package main

import "fmt"

func main() {
	price := 4.50 // float64

	itemQnty := 13 // int

	// Go doesn't convert types automatically
	totalIncome := price * float64(itemQnty)

	fmt.Printf("Total income during a day %.2f\n", totalIncome)
}
