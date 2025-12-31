package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Customer struct {
	Name string
}

type Staff struct {
	Role string
}

// implement Greet method for Greeter interface
func (c Customer) Greet() string {
	return fmt.Sprintf("Customer %s says: Hello!", c.Name)
}

func (s Staff) Greet() string {
	return fmt.Sprintf("Staff %s says: Hello! Welcome to the club, buddy :)", s.Role)
}

func main() {
	greeters := []Greeter{
		Customer{Name: "Omar"},
		Staff{Role: "Barista"},
		Customer{Name: "Maria"},
	}

	for _, greeter := range greeters {
		fmt.Println(greeter.Greet())
	}

	fmt.Println()

	greeters = append(greeters, Staff{Role: "Cleaner"})

	for _, greeter := range greeters {
		fmt.Println(greeter.Greet())
	}
}
