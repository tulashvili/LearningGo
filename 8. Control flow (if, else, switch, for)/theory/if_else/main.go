package main

import "fmt"

func main() {
	customerName := "Omar"
	orderedSize := "Large"

	if orderedSize == "Small" {
		fmt.Println(customerName, "order done")
	} else {
		fmt.Println(customerName, "order not done. We need more time")
	}
}
