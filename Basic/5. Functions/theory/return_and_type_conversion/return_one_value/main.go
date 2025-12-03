package main

import "fmt"

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2) // требуется явное преобразование типов
	return loyaltyPoints
}

func updateTotalPoints(currentPoints int, newPoints int) int {
	return currentPoints + newPoints
}
func main() {
	totalPoints := 120
	newlyEarnedPointsToday := calculateLoyaltyPoints(9.49) // 18
	fmt.Println("Earned points today:", newlyEarnedPointsToday)

	totalPoints = updateTotalPoints(totalPoints, newlyEarnedPointsToday)
	fmt.Println("Updated total points today:", totalPoints) // 120 + 18
}
