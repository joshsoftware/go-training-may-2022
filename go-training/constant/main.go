package main

import (
	"fmt"
	"math"
)

func main() {

	const Pi float32 = 3.141592653589793238
	fmt.Println(Pi)

	const num = 10
	fmt.Println(num)

	const (
		a int = 1
		b     = 1.414
		c     = "hello"
	)
	fmt.Println(a, b, c)

	// const m = math.Sqrt(4)  //is not allowed
	var m = math.Sqrt(4) // allowed
	fmt.Println(m)

}
