package main

import (
	"fmt"
	"strings"
)

type MenuList []string

func (ml MenuList) String() string {
	// [Tea, Cofee, Water]
	// OPTION 1
	return fmt.Sprintf("[%s]", strings.Join(ml, ", "))

	// OPTION 2
	// c := "["
	// for i, menuItem := range ml {
	// 	c += menuItem
	// 	if i < len(ml)-1 {
	// 		c += ", "
	// 	}
	// }
	// c += "]"

	// return c

	// OPTION 3
	// return "[" + strings.Join(ml, ", ") + "]"
}
func main() {
	order := MenuList{"Tea", "Cofee", "Water"}
	fmt.Println(order)
}
