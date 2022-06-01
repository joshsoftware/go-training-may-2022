package main

import "fmt"

func main() {
	a := 3
	var ptr1 *int
	ptr1 = &a
	fmt.Println(*ptr1, "adress iss ", ptr1)
}
