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
func main() {
	user := User{
		Name:   "Omar",
		Age:    26,
		Rating: 20.7,
	}
	user.GoodBye()
	user.Greeting()
}
