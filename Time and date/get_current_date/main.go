package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Printf("The time is %d %s %d\n", currentTime.Day(), currentTime.Month(), currentTime.Year())
}
