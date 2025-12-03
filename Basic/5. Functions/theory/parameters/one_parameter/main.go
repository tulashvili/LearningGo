package main

import "fmt"

func greet(name string) {
	var greetMsg = "Welcome to the coffee shop"
	fmt.Println(greetMsg, name)

}
func main() {
	greet("Negroni")
	greet("So So Cofee")
}
