package main

import "fmt"

func main() {
	var points int

	fmt.Println("Print your points:")
	fmt.Scan(&points)
	fmt.Printf("You have %d points\n", points)

	members := []string{"free", "premium", "platinum"}

	if points >= 10 && points < 40 {
		fmt.Printf("You are %s member\n", members[0])
	} else if points >= 40 && points < 60 {
		fmt.Printf("You are %s member\n", members[1])
	} else if points >= 60 {
		fmt.Printf("You are %s member\n", members[2])
	} else {
		fmt.Println("Huy znaet, you are bobrovnikov :)")
	}
}
