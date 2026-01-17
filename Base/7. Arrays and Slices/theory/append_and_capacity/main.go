package main

import "fmt"

func main() {
	menu := []string{"Cake", "Pie"}
	fmt.Println("Initial menu:", menu)
	fmt.Println("Lenght old menu slice:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Slice memory location:%p\n", &menu)
	fmt.Printf("First element slice mem loc: %p\n", &menu[0])
	fmt.Printf("Second element slice mem loc: %p\n", &menu[1])

	fmt.Println("------------------")

	menu = append(menu, "Shawerma")
	fmt.Println("New menu:", menu)
	fmt.Println("Lenght new menu slice:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Slice memory location:%p\n", &menu)
	fmt.Printf("First element slice mem loc: %p\n", &menu[0])
	fmt.Printf("Second element slice mem loc: %p\n", &menu[1])

	fmt.Println("------------------")

	menu = append(menu, "Burger")
	fmt.Println("New menu:", menu)
	fmt.Println("Lenght new menu slice:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Slice memory location:%p\n", &menu)
	fmt.Printf("First element slice mem loc: %p\n", &menu[0])
	fmt.Printf("Second element slice mem loc: %p\n", &menu[1])

	fmt.Println("------------------")

	menu = append(menu, "Ice cream")
	fmt.Println("New menu:", menu)
	fmt.Println("Lenght new menu slice:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Slice memory location:%p\n", &menu)
	fmt.Printf("First element slice mem loc: %p\n", &menu[0])
	fmt.Printf("Second element slice mem loc: %p\n", &menu[1])

	fmt.Println("------------------")

	menu = append(menu, "Orange", "Banana", "Apple", "Pineapple")
	fmt.Println("New menu:", menu)
	fmt.Println("Lenght new menu slice:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Slice memory location:%p\n", &menu)
	fmt.Printf("First element slice mem loc: %p\n", &menu[0])
	fmt.Printf("Second element slice mem loc: %p\n", &menu[1])

	fmt.Println()

	cupSizes := make([]string, 0, 4)
	fmt.Println(cupSizes)
	fmt.Println("Len cupSizes:", len(cupSizes), "Cap cupSizes:", cap(cupSizes))
	// cupSizes[0] = "Medium" // panic: runtime error: index out of range [0] with length 0

	cupSizes = append(cupSizes, "Medium", "Medium2", "Medium3", "Medium4", "Medium5") // capacities are auto scaling
	cupSizes[0] = "Extra medium"
	fmt.Println("Len cupSizes:", len(cupSizes), "Cap cupSizes:", cap(cupSizes))
	fmt.Println(cupSizes)

}
