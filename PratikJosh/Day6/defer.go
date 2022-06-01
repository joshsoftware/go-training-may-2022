package main

import "fmt"

func main() {
	i := 23
	defer fmt.Println("I=", i)
	i = 33
	defer fmt.Println("I=", i)
	foo()
}
func foo() {

	a := 4
	b := 6
	defer func() {
		fmt.Println(a + b)
	}()
	defer fmt.Println(a + b)
	return
}
