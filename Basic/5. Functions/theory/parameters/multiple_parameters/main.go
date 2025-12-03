package main

import "fmt"

func getDrinkInfo(customerName string, drink string) {
	fmt.Printf("%s's favourite drink is %s\n", customerName, drink)

}

func main() {
	getDrinkInfo("Omar", "Milk")
}
