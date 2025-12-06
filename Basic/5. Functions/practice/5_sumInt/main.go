package main

import "fmt"

func sumInt(a ...interface{}) {

	var sum int
	var count int

	for _, elem := range a {
		num, ok := elem.(int)

		if ok {
			sum += num
			count++
		} else {
			continue
		}

	}
	fmt.Println(count, sum)
}
func main() {
	var firstNum int
	var secondNum int

	fmt.Scan(&firstNum)
	fmt.Scan(&secondNum)

	sumInt(firstNum, "hello", secondNum)

}
