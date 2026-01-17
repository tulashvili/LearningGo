package main

import "fmt"

type Barista interface {
	PrepareCoffee() string
}

type SeniorBarista struct {
	Name string
}

type JuniorBarista struct {
	Name string
}

// приготовление кофе
func (s SeniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s made a latte", s.Name)
}

func (j JuniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s made a hot chocolate", j.Name)
}

// подача кофе клиенту
func ServeDrink(b Barista) {
	fmt.Println(b.PrepareCoffee())
	fmt.Println("Barista served coffee to the client")
	fmt.Println()
}

func main() {
	firstBarista := SeniorBarista{Name: "Omar"}

	// даем возможность переопределять secondBarista
	var secondBarista Barista
	secondBarista = JuniorBarista{Name: "Julia"}

	// fmt.Println(firstBarista.PrepareCoffee())
	// fmt.Println(secondBarista.PrepareCoffee())

	// Omar made a latte
	// Barista served coffee to the client
	ServeDrink(firstBarista)

	// 	Julia made a hot chocolate
	// Barista served coffee to the client
	ServeDrink(secondBarista)

	fmt.Println("Переопределяем secondBarista")
	secondBarista = SeniorBarista{Name: "Julia"}
	ServeDrink(secondBarista)

}
