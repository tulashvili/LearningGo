package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("My favourite number is", rand.Intn(2))
	color.Cyan("My favourite number is", rand.Intn(2))
}
