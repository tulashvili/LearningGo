package main

import (
	"fmt"
)

func main() {
	// days := []string{"Monday", "Tuesday", "Wednesday"}
	// today := "Sunday"
	// found := false

	// for _, day := range days {
	// 	if day == today {
	// 		found = true
	// 		break
	// 	}
	// }
	// // findDay := slices.Contains(days, "Sunday")
	// if found {
	// 	fmt.Println("Cool")
	// } else {
	// 	fmt.Println("Chto ti, tvar?")
	// }

	today := "Wendesday"
	switch today {
	case "Monday", "Tuesday":
		fmt.Println("Start a week")
	case "Friday":
		fmt.Println("Finish a week")
	default:
		fmt.Println("Middle a week")
	}
}
