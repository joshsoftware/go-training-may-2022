package main

import (
	//when we want to import any package then we need to give a path of that package here
	//we can also give a alias name to that package
	//in this example calc is alias name for below imported package
	calc "GoDilyCodes/go-training-may-2022/RohitSangamnerkar/day3/fonctions/calculator"
	"fmt"
)

func main() {
	var a int = 100
	var b int = 10

	//now call functions from other file

	fmt.Println("Addition of two numbers : ", calc.Sum(a, b))
	fmt.Println("Subtraction of  two numbers : ", calc.Sub(a, b))
	fmt.Println("Multiplication of two number : ", calc.Mul(a, b))
	fmt.Println("Division of two numbers : ", calc.Div(a, b))

	//we can also create a anonymus function
	//a function which has no name

	//lets see one example

	func() {
		fmt.Println("Testing anonymous function")
	}() //here it will call this function

}
