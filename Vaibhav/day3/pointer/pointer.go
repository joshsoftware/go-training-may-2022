package main

import "fmt"

func main() {
	var a int = 20
	var i *int // pointer variable declaration

	i = &a // store address of a in pointer variable

	fmt.Printf("Address of a variable: %x\n", &a)

	// address stored in pointer variable
	fmt.Printf("Address stored in i variable: %x\n", i)

	// accessing the value using the pointer
	fmt.Printf("Value of *i variable: %d\n", *i)
}
