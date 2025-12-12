package main

import "fmt"

func main() {
	var potatoSize = 50
	fmt.Println("potato value:", potatoSize)
	fmt.Println("potato address memory:", &potatoSize)

	var potatoSizePointer *int = new(int)
	potatoSizePointer = &potatoSize
	fmt.Println("potatoSizePointer address potato", potatoSizePointer)
	fmt.Println("potatoSizePointer value", *potatoSizePointer)
	fmt.Println("potatoSizePointer address", &potatoSizePointer)

	// potato = 50
	// potatoSizePointer -> potato [50] = 0xc000012050
	// potatoSizePointer = 50
	// &potatoSizePointer = 0xc00009c030
}
