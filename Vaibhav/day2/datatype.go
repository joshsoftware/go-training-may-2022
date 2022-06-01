package main

import "fmt"

func main() {
	a := 23                  //int
	var b string = "vaibhav" //string
	var c float64 = 23.23    //float
	var d bool = true        // boolean

	fmt.Println(a, b, c, d)
	fmt.Printf("value of A is %v  %d and datatype is %T\n", a, a, a)
	fmt.Printf("value of B is %s  %s and datatype is %T\n", b, b, b)
	fmt.Printf("value of c is %f %v and datatype is %T\n", c, c, c)
	fmt.Printf("value of D is %v  %t and datatype is %T\n", d, d, d)
}
