package main

import "fmt"

func main() {
	menu := [3]string{"Tea", "Coffee", "Juice"}
	slice := menu[:2]

	fmt.Println("Before slice change:")
	fmt.Println("menu:", menu)
	fmt.Println("slice:", slice)

	slice[0] = "Water"
	fmt.Println("After slice change:")
	fmt.Println("menu:", menu)
	fmt.Println("slice:", slice)
}
