package main

import "fmt"

func main() {
	var coffeeSizes [3]string
	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "Small"
	fmt.Println(coffeeSizes)

	coffeeSizes[1] = "Medium"
	coffeeSizes[2] = "Large"
	fmt.Println(coffeeSizes)

	coffeeSizes[2] = "Extra Large"
	fmt.Println(coffeeSizes)
	fmt.Println("First element is:", coffeeSizes[0])

	// coffeeSize[1] = true // compile error
	// coffeeSize[5] compile error
}
