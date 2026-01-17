package main

import "fmt"

type AnyValue interface{}

func LogAnyValue(v interface{}) {
	fmt.Println(v)

}
// any is equivalent interface{}
func LogAnyValueWithAny(v any) {
	fmt.Println(v)

}
func main() {
	// can assign value of any type
	var any AnyValue = "some string"
	fmt.Println(any)

	any = 10
	fmt.Println(any)

	var anotherAny interface{} = "Latte"
	fmt.Println(anotherAny)

	// slice accepts values of any types
	diffTypesValues := []interface{}{
		"Latte",
		50.5,
		true,
		[3]int{10, 5, 3},
	}

	fmt.Println(diffTypesValues...)

	// call a func with any value
	LogAnyValue("Omar")
	LogAnyValue(true)
	LogAnyValue([3]string{"Omar", "Tulashvili"})
}
