package main

import "fmt"

func main() {
	drinks := []string{"Latte", "Espresso", "Expired Milk", "Capuccino", "Water"}

	for _, drink := range drinks {
		if drink == "Expired Milk" {
			fmt.Println("Skip expired drink:", drink)
			continue
		}
		if drink == "Capuccino" {
			fmt.Printf("The %s is out, please choose another one\n", drink)
			break
		}
		fmt.Println("Preparing drink:", drink)
	}
}
