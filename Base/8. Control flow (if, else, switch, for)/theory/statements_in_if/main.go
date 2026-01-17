package main

import "fmt"

func getFullOrderAmount(value float64, tax float64) float64 {
	return value + value*tax
}
func main() {
	// var points int

	// fmt.Println("Print points (int)")
	// fmt.Scan(&points)

	if points := 16; points > 15 {
		fmt.Println("You have dohera points")
	}

	if fullAmount := getFullOrderAmount(14.50, 0.1); fullAmount > 15 {
		fmt.Printf("Your order amount is %.2f. You live in Europe, because pisya bez glaz", fullAmount)
	}

	totalLoyaltyPoint := 100

}
