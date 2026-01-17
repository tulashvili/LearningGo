package main

import "fmt"

// describe structure
type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func main() {
	// create, initialize structure
	user := User{
		Name:        "Omar",
		Age:         26,
		PhoneNumber: "+7 999 888 77 66",
		IsClose:     false,
		Rating:      2.5,
	}

	fmt.Println("User:", user)
	fmt.Println("Name before:", user.Name)
	user.Name = "Jacob"
	fmt.Println("Name after:", user.Name)
}
