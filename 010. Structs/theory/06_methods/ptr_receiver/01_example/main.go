package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Rating float64 // expect: 0...10
}

func (u User) Greeting() {
	fmt.Println("Hi all")
	fmt.Println("My name is", u.Name)
	fmt.Printf("I have %d years old\n", u.Age)
	fmt.Println("My rating is:", u.Rating)
	fmt.Println()

}

func (u User) GoodBye() {
	fmt.Println("Goodbye all")
	fmt.Println("My name was", u.Name)
	fmt.Printf("I had %d years old\n", u.Age)
	fmt.Println("My rating was:", u.Rating)
	fmt.Println()
}

// add validation, that rating have expected value
func (u *User) UpRating(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println("I added rating value")
	} else {
		fmt.Println("It doesn't pass validation")
	}
}
func main() {
	user := User{
		Name:   "Omar",
		Age:    26,
		Rating: 4.7,
	}
	user.GoodBye()
	user.Greeting()

	ptr := &user
	fmt.Println("Rating before UP:", ptr.Rating)
	user.UpRating(3)
	fmt.Println("Rating after UP:", ptr.Rating)

}
