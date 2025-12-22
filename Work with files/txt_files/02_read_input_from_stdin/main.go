package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // map, который хранит количество раз (int), которое встретилась string
	input := bufio.NewScanner(os.Stdin) // читает данные из os.Stdin. по умолчанию читает по строкам (разделитель — \n)

	// Выполняем Scan читает введенные строки - выполняем тело
	for input.Scan() {
		line := input.Text()
		counts[input.Text()]++

		fmt.Printf("%s встретилась - %d раз\n", line, counts[line])

		if counts[line] > 3 {
			fmt.Printf("%d\t %s\n", counts[line], line) // 	// Если строка встретилась больше 3 раз — печатаем её
		}
	}
	// Из цикла никогда не выйдет, пока не сделаем break
}
