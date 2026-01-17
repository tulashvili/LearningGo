package main

import "fmt"

func main() {
	dailyIncome := []float64{127.23, 100.10, 54.20, 23.12, 80.26, 44.11, 76.89}
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println("Weekly income report:")

	var totalIncome float64
	for i, income := range dailyIncome {
		day := daysOfWeek[i]
		fmt.Printf("	%s income is $%.2f\n", day, income)
		totalIncome += income
	}
	fmt.Printf("\nTotal weekly income: $%.2f\n", totalIncome)
}
