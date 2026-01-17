package main

import "fmt"

func main() {
	var hour int
	var isMember bool
	var orderAmount float64

	fmt.Println("Print hour (int)")
	fmt.Scan(&hour)

	fmt.Println("Print membership (true or false)")
	fmt.Scan(&isMember)

	fmt.Println("Print order amount (float 64)")
	fmt.Scan(&orderAmount)

	// RECOMEND!
	if hour >= 15 && hour <= 17 && isMember && orderAmount > 10 {
		fmt.Println("Get your free americano!")
	} else {
		fmt.Println("We don't get americano to you")
	}
	// NOT RECOMENDED!
	if hour >= 15 {
		if hour <= 17 {
			if isMember {
				if orderAmount > 10 {
					fmt.Println("Get your free americano!")
				}
			}
		}
	}
}
