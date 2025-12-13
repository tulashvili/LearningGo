package main

import "fmt"

func main() {
	var numbersOfNumbers int // Считаем количество чисел
	fmt.Println("Введите кол-во чисел в последовательности:")
	fmt.Scan(&numbersOfNumbers)

	var allNumbers int // Вводим числа
	fmt.Printf("Введите %d чисел:", numbersOfNumbers)
	fmt.Scan(&allNumbers) // Считываем числа в последовательности

	sum := 0

	for i := 0; i < numbersOfNumbers; i++ {

		if (allNumbers >= 10 && allNumbers <= 99) && (allNumbers%8 == 0) {
			sum += allNumbers
		}
	}
	fmt.Println(sum)
}
