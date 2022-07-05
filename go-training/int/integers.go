package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int = 10
	b := 20
	c := -30

	fmt.Println("values of a : ", a, "values of b : ", b, "values of c : ", c)
	fmt.Printf("Type of a is %T,Size of a is %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("Type of b is %T,Size of b is %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("Type of c is %T,Size of c is %d\n", c, unsafe.Sizeof(c))

}
