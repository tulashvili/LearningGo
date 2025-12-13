package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите значение 'a'")
	fmt.Scan(&a)

	fmt.Println("Введите значение 'b'")
	fmt.Scan(&b)

	fmt.Println()
	iterationNum := 0

	var sum int

	fmt.Println("Значение 'a' до цикла = ", a)
	fmt.Println("Значение 'a' до цикла = ", b)
	fmt.Println("Значение 'sum' до цикла = ", sum)
	fmt.Println()
	fmt.Printf("Выполняем сложение диапазона значений от %d до %d\n", a, b)

	for i := a; i <= b; i++ {
		iterationNum++
		sum += i
		fmt.Printf("#%d: a = %d, i = %d, sum = %d\n", iterationNum, a, i, sum)
	}
	fmt.Println()
	fmt.Println("Значение 'a'после цикла = ", a)
	fmt.Println("Значение 'a' после цикла = ", b)
	fmt.Println("Значение 'sum' после цикла = ", sum)
}
