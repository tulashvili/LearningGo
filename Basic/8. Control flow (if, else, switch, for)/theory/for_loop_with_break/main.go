package main

import "fmt"

func main() {
	fmt.Println("Welcome to the club, buddy")

	for {
		var person string
		fmt.Println("Enter your inviter (or type 'bye' to quit):")
		fmt.Scanln(&person)

		if person == "bye" {
			fmt.Println("Nu i nahyu ti pripersya to our club?")
			break
		}

		if person == "" {
			fmt.Println("I need your inviter name, suchka!")
			continue
		}

		fmt.Println("Bodyguard following you to inviter - ", person)
	}

	fmt.Println("Finishing story ....")
}
