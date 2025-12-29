package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // map, который хранит количество раз (int), которое встретилась string
	files := os.Args[0:]           // берем аргументы командной строки

	if len(files) == 0 { // если файлы не были переданы, то читаем с standart input (ввод)
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files { // если файлы передали, обрабатываем каждый
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
