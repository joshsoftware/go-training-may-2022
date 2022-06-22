package main

import "fmt"

func main() {

	/*	In Go we have 4 types of Operators are there, They are

		1. Arithmetic Operators
		2. Assignment Operators
		3. Relational Operators
		4. Logical Operators */

	// 	1. Arithmetic Operators :
	// + : Addition

	var numOne, numTwo = 10, 20
	fmt.Println("+ Operator :", numOne+numTwo)

	// - : substraction
	fmt.Println("- Operator :", numOne-numTwo)

	// * : multiplication
	fmt.Println("* Operator :", numOne*numTwo)

	// / : Divison
	fmt.Println("/ Operator :", numTwo/numOne)

	// % : Modulo Division
	fmt.Println("% Operator :", numTwo%numOne)

	// actual result from divison
	a := 11.0
	b := 4.0
	div := a / b
	fmt.Printf("%g / %g = %g", a, b, div)

}
