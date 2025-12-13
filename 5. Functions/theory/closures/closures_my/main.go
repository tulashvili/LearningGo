package main

import "fmt"

// На текущий момент мы не можем, без знания темы вызова функций (return func())
// и явного определения значения var employmentSalary float64 = 3000 решить вопрос
func createSalaryRegulator() (func(change float64) float64, float64) {
	var employmentSalary float64
	salaryRegulator := func(change float64) float64 {
		employmentSalary = employmentSalary + change
		return employmentSalary
	}
	return salaryRegulator, employmentSalary
}

func main() {
	changedSalary, employmentSalary := createSalaryRegulator()
	fmt.Printf("My employment salary this month in %.2f\n", employmentSalary)
	fmt.Printf("My salary changed on --- %.2f USD ---\n", changedSalary(2500))
	fmt.Printf("My salary changed on --- %.2f USD ---\n", changedSalary(500))

}
