package main

import "fmt"

type CoffeeMachine interface {
	Brew() string
	// Clean() string
}

// capsule machine implementation
type CapsuleMachine struct {
	Brand string
	Model string
	Price int
}
// drip machine implementation
type DripMachine struct {
	Brand string
	Model string
	Price int
}

// brew - go видит таким образом реализацию интерфейса
func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s %s brewed a cup of capsule coffee", c.Brand, c.Model)
}

func (c CapsuleMachine) Clean() string {
	return fmt.Sprintf("Machine is cleaning: %s %s", c.Brand, c.Model)
}

func (d DripMachine) Brew() string {
	return fmt.Sprintf("%s %s brewed a cup of drip coffee", d.Brand, d.Model)
}

func (d DripMachine) Clean() string {
	return fmt.Sprintf("Machine is cleaning: %s %s", d.Brand, d.Model)
}

func main() {
	var capsuleMachine CoffeeMachine
	var dripMachine CoffeeMachine

	capsuleMachine = CapsuleMachine{
		Brand: "Nespresso",
		Model: "AB431",
		Price: 135,
	}
	dripMachine = DripMachine{
		Brand: "Philips",
		Model: "HB342",
		Price: 185,
	}

	// мы не можем вызвать метод clean,
	// потому что capsuleMachine и dripMachine имеет тип CapsuleMachine
	// а тип CoffeeMachine не имеет метода clean
	fmt.Println(capsuleMachine.Brew())
	// fmt.Println(capsuleMachine.Clean())

	fmt.Println(dripMachine.Brew())
	// fmt.Println(dripMachine.Clean())

	// мы можем использовать метод clean для anotherDripMachine
	// потому что переменная anotherDripMachine имеет тип DripMachine
	// а тип DripMachine имеет реализацию метода Clean
	anotherDripMachine := DripMachine{
		Brand: "SuperDrip",
		Model: "FC561",
	}

	fmt.Println(anotherDripMachine.Clean())

}
