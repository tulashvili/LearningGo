package main

import "fmt"

type Action struct {
	On bool

	// Боеприпасы персонажа
	Ammo int

	// Ресурс мотоцикла
	Power int
}

// Выстрел можно делать только при наличии Ammo
func (a *Action) Shoot() bool {
	if a.Ammo > 0 {
		a.Ammo--
		a.On = true
		fmt.Println("Shoot done")
		fmt.Println("ammunition left:", a.Ammo)
	} else {
		fmt.Println("Shoot doesn't do.")
		fmt.Println("Not enough ammo")
		a.On = false
	}
	fmt.Println()
	return a.On
}

func (a *Action) RideBike() bool {
	if a.Power > 0 {
		a.Power--
		a.On = true
		fmt.Println("you got on a motorcycle")
		fmt.Println("power left:", a.Power)
	} else {
		fmt.Println("you can't get on a motorcycle")
		fmt.Println("Not enough Power")
		a.On = false
	}
	fmt.Println()
	return a.On
}
func main() {
	action := Action{
		On:    true,
		Ammo:  0,
		Power: 10,
	}

	testStruct := &action
	action.Shoot()
	action.RideBike()
	fmt.Println(testStruct)

}
