package main

import "fmt"

// каждая кофемашина должна уметь в эти методы
// и не важно, что это за кофемашина
type CoffeeMachine interface {
	Brew() string
}

// капусульная кофемашина имеет такие поля
type CapsuleMachine struct {
	Brand string
}

type FrenchPressMachine struct {
	Brand string
}

// создаем метод для типа CapsuleMachine
// функция называется именно так, потому что мы реализуем метод Brew из CoffeeMachine
func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s brewed one cup of Coffee", c.Brand)
}

// создаем метод для типа FrenchPressMachine
// функция называется именно так, потому что мы реализуем метод Brew из CoffeeMachine
func (f FrenchPressMachine) Brew() string {
	return fmt.Sprintf("%s brewed one cup of Coffee", f.Brand)
}

func main() {
	var capsuleMachine CoffeeMachine
	capsuleMachine = CapsuleMachine{Brand: "Nespresso"}
	fmt.Println(capsuleMachine.Brew())

	var frenchPressMachine CoffeeMachine
	frenchPressMachine = FrenchPressMachine{Brand: "Philips"}
	fmt.Println(frenchPressMachine.Brew())

}
