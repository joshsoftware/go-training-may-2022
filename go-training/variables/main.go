package main

import "fmt"

func main() {

	var name string
	name = "Rak"

	num := 8

	var name1, age, id = "Raj", 26, 1
	fmt.Println(name)
	fmt.Println(num)
	fmt.Println(name1, age, id)

	//block level declaration
	var (
		a int    = 1
		b int    = 2
		c string = "hello"
	)
	fmt.Println(a, b, c)

	// default values:
	var x int
	var y string
	var z bool

	fmt.Println("x :", x, "y :", y, "z :", z)
}
