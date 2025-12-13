package main

import "fmt"

//	func estimateCookingTime(cupsQty int, secondsPerCup int) int {
//		totalTimeSeconds := cupsQty * secondsPerCup
//		return totalTimeSeconds
//	}
func estimateCookingTime(cupsQty int, secondsPerCup int) (totalTimeSeconds int, measure string) {
	totalTimeSeconds = cupsQty * secondsPerCup
	measure = "seconds"
	return // naked return
}
func main() {
	brewTime, measure := estimateCookingTime(10, 60)
	fmt.Printf("Estimated brew time: %d %s\n", brewTime, measure)
}
