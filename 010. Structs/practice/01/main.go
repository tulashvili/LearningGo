package main

import "fmt"

type Action struct {
	On bool

	// Боеприпасы персонажа
	Ammo int

	// Ресурс мотоцикла
	Power int
}

// Оба метода срабатывают при условии, что action.On = True

func (a *Action) Shoot() bool {
	if a.On != true {
		fmt.Println("a.On = false, shoot false")
	} else if a.Ammo > 0 {
		a.Ammo--
		fmt.Println("Shoot true")
	} else {
		fmt.Println("Shoot false")
		fmt.Println("Ammo:", a.Ammo)

	}
	fmt.Println()
	return a.On
}

func (a *Action) RideBike() bool {
	if a.On != true {
		fmt.Println("a.On = false, RideBike false")
	} else if a.Power > 0 {
		a.Power--
		fmt.Println("RideBike true")
	} else {
		fmt.Println("RideBike false")
		fmt.Println("Power:", a.Power)
	}
	fmt.Println()
	return a.On
}

func main() {
	action := Action{
		On:    true,
		Ammo:  0,
		Power: 2,
	}

	testStruct := &action
	testStruct.Shoot()
	testStruct.RideBike()
}
