package main

import "fmt"

func main() {
	coffeeBars := map[string]string{
		"Lenina":   "Have",
		"Pushkina": "Have",
		"Mira":     "Have",
	}
	bar := "Mira"

	street, exists := coffeeBars[bar]
	fmt.Println("Exists:", exists)
	
	if exists {
		fmt.Println("Street:", street)
	} else {
		fmt.Printf("%s is not in Coffee Bars list", bar)
	}
}
