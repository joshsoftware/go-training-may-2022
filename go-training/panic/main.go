package main

import "fmt"

func Div(a, b int) {
	if b == 0 {
		panic("Cannot devide a number by zero.")
	} else {
		result := a / b
		fmt.Println("Result :", result)
	}
}

func main() {

	Div(6, 3)
	Div(6, 0)
	Div(2, 8)

}
