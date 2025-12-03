package main

import "fmt"

func createTempRegulator() (func(change float64) float64, float64) {
	baseTemp := 80.0

	tempRegulator := func(change float64) float64 {
		baseTemp = baseTemp + change
		return baseTemp
	}
	return tempRegulator, baseTemp
}
func main() {
	// get tempRegulator, baseTemp from return
	adjustRegulator, originalTempreture := createTempRegulator()
	// tempRegulator := adjustRegulator(-5.7)

	fmt.Printf("orig temp is: %.1f degrees\n", originalTempreture)
	fmt.Printf("currently temp is: %.1f degrees\n", adjustRegulator(-5.7))
	fmt.Printf("currently temp is: %.1f degrees\n", adjustRegulator(-5.7))

	// dynamicTempreture := tempRegulator - originalTempreture
	// fmt.Printf("tempreture changed by %.1f\n", dynamicTempreture)

}
