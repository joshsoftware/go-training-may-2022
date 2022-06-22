package main

import "fmt"

//recover function for handling panic
func HandlePanic() {

	//detect if panic occurs or not
	a := recover()
	if a != nil {
		fmt.Println("Recover :", a)
	}
}

func Div(a, b int) {

	//execute the HandlePanic even after panic occurs
	defer HandlePanic()

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
