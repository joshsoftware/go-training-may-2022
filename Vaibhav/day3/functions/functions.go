package main

import "fmt"

//Defining a Function

// func function_name( [parameter list] ) [return_types]
// {
// body of the function
// }

// function returning the max between two numbers
func Max(num1, num2 int) int {

	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//Returning multiple value from Function
func Swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("maximum no is ", Max(30, 20))

	a, b := Swap("Vaibhav", "Gholap")
	fmt.Println(a, b)

}
