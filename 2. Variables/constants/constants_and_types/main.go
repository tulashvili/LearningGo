package main

import "fmt"

func main() {
	// untyped constant with integer value
	const rewardPoints = 10
	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints) // int

	// adding untyped constant to a float64 - valid, constant adapts
	var totalRewardPoints float64 = 150.3
	totalRewardPoints += rewardPoints

	fmt.Printf("Updated loyalty points %.2f\n", totalRewardPoints)
}
