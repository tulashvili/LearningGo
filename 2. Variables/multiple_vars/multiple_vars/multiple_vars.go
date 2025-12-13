package multiple_vars

import "fmt"

func multiple_vars() {
	cofeeName, size, price := "Espresso", "Medium", 2.50
	// var size = "Medium"
	// var price = 2.50

	fmt.Println("Medium Espresso price is $2.50")
	fmt.Println(size, cofeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, cofeeName, price)
}
