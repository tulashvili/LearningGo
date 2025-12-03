package main

import "fmt"

func main() {
	var (
		coffeType = "Latte"
		quantity  = 8
		unitPrice = 4.25
	)

	const (
		SizeSmall  = "S"
		SizeMedium = "M"
		SizeLarge  = "L"
	)
	fmt.Printf("Ordered %d %s %s size with prices $%.2f\n", quantity, coffeType, SizeLarge, unitPrice)
}
