package main

import "fmt"

func sumInt(array ...interface{}) (a int, b int) {

	var sum int
	var count int

	for _, elem := range array {
		num, ok := elem.(int)

		if ok { // check, that elem is int
			sum += num
			count++
		} else {
			continue
		}

	}
	return count, sum
}
func main() {
	a, b := sumInt(1, 0)
	fmt.Printf("%d, %d\n", a, b)

}
