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
	if a.On && a.Ammo > 0 {
		a.Ammo--
		return true
	}
	return false
}

func (a *Action) RideBike() bool {
	if a.On && a.Power > 0 {
		a.Power--
		return true
	}
	return false
}

func main() {
	action := Action{
		On:    true,
		Ammo:  1,
		Power: 1,
	}

	testStruct := &action
	testStruct.Shoot()
	testStruct.RideBike()
}
