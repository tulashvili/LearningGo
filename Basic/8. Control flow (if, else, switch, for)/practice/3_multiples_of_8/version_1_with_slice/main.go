package main

import "fmt"

func main() {
	// numOfNums := 5
	nums := []int{38, 24, 800, 8, 16}
	multNum := 8
	multNums := []int{}
	var iterationNum int

	for _, v := range nums {
		iterationNum++
		fmt.Printf("#%d - %d\n", iterationNum, v)

		if v > 10 && v < 100 && v%multNum == 0 {
			fmt.Printf("Число %d двузначное и делится на %d без остатка\n", v, multNum)
			multNums = append(multNums, v)
		}
		fmt.Println()
	}
	fmt.Printf("Двузначные числа, что делятся на %d без остатка %d\n", multNum, multNums)

	var sumNums int
	for i := 0; i < len(multNums); i++ {
		sumNums += multNums[i]
	}
	fmt.Println(sumNums)

}
