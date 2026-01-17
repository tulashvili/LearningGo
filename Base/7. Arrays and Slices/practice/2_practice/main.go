package main

import "fmt"

func main() {
	array := [5]int{} // create array with 5 num - 0,0,0,0,0
	var inputNum int  // save input num

	for index := 0; index < 5; index++ {
		fmt.Scan(&inputNum)     // input num
		array[index] = inputNum // save user input num to array in position [index]
		fmt.Println(array)
	}
	maxNum := array[0]
	for index := 0; index < len(array); index++ { // get first element from array
		if array[index] > maxNum {
			maxNum = array[index]
		}
	}
	fmt.Println("Самое большое число = ", maxNum)
}
