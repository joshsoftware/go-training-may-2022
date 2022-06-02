package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func mult(a float32, b float32) float32 {
	return a * b
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(mult(10.5, 20.5))
}
