package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

//named return values :
func mul(a, b int) (result int) {
	result = a * b
	return
}

//multiple return values :
func calc(a, b int) (int, int) {
	c := a + b
	d := a * b
	return c, d
}

//variadic functions:
func variadic(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// //anonymus functions:
// s:=func(a,b int){
// 	fmt.Println(a+b)
// }(6,7)

//main function
func main() {
	a := 6
	b := 8
	res := sum(a, b)
	fmt.Println("a + b =", res)
	fmt.Println("a * b =", mul(a, b))
	//m-r
	x, y := calc(3, 18)
	fmt.Println("multi-return :", x, y)
	//variadic
	variadic(1, 2, 3, 4)

	//anonymus functions:
	s := func(a, b int) {
		fmt.Println(a + b)
	}
	s(6, 7)
}
