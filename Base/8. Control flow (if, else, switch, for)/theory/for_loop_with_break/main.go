package main

import "fmt"

func main() {
	fmt.Println("Welcome to the club, buddy")

	for {
		var person string
		fmt.Println("Enter your inviter (or type 'bye' to quit):")
		fmt.Scanln(&person)

		if person == "bye" {
			fmt.Println("What do you doing there?")
			break
		}

		if person == "" {
			fmt.Println("I need your inviter name!")
			continue
		}

		fmt.Println("Bodyguard following you to inviter - ", person)
	}

	fmt.Println("Finishing story ....")
}
