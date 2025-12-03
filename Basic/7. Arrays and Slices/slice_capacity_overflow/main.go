package main

import "fmt"

func main() {
	desserts := [3]string{"Cake", "Pie", "Burger"}
	fmt.Println("Initial array:", desserts)
	fmt.Println("Len of array:", len(desserts), "Cap of array:", cap(desserts))

	fmt.Println()

	slice := desserts[:1]
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice))
	fmt.Println("Array after \"desserts[:1]\" (its slice):", slice)

	fmt.Println()

	slice = append(slice, "Macachino")
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice))
	fmt.Println("Updated slice:", slice)
	fmt.Println("Creates new array:", desserts)

	fmt.Println()
	// Because len is already equal to cap ->
	// old array not communicate with slice,
	// but new array is allocated (comunicate) for slice
	slice = append(slice, "Ice cream", "Macachino")
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice))
	fmt.Println("Updated slice:", slice)
	fmt.Println("Array is not create:", desserts)

	fmt.Println()

	slice[0] = "Choco" // original array is not modified
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice))
	fmt.Println("Updated slice:", slice)
	fmt.Println("Array is not create:", desserts)

}
