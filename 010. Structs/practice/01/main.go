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
	if a.Ammo > 0 && a.On {
		a.Ammo--
		fmt.Println("Shoot done")
		fmt.Println("ammunition left:", a.Ammo)
	} else {
		fmt.Println("Shoot doesn't do.")
		fmt.Println("Ammo:", a.Ammo)
		fmt.Println("On:", a.On)
		a.On = false
	}
	fmt.Println()
	return a.On
}

func (a *Action) RideBike() bool {
	if a.Power > 0 && a.On {
		a.Power--
		fmt.Println("you got on a motorcycle")
		fmt.Println("power left:", a.Power)
	} else {
		fmt.Println("You can't got on a motorcycle")
		fmt.Println("Ammo:", a.Power)
		fmt.Println("On:", a.On)
	}
	fmt.Println()
	return a.On
}
func main() {
	action := Action{
		On:    false,
		Ammo:  2,
		Power: 10,
	}

	testStruct := &action
	action.Shoot()
	action.RideBike()
	fmt.Println(testStruct)

}
