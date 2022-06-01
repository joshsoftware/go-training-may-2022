package main

import (
	"fmt"
)

func printInterface(x interface{}) {
	fmt.Printf("%v is of type %T \n", x, x)
}

type User struct {
	FirstName string
	LastName  string
}

// Stringer Interface

func (u User) String() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func learnTypeAssertion1() {
	// Now the x type will be string
	var x interface{} = "Shubham"
	// Skipping panic handling
	var s string = x.(string)
	fmt.Println(s)

	// It wont panic as we are asserting the type of actual interface
	str, ok := x.(string)
	fmt.Println(str, ok)

	i, ok := x.(int)
	fmt.Println(i, ok)

	var in int = x.(int)
	fmt.Println(in)
}

func learnTypeAssertion2() {
	var x interface{} = "Shubham"

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is of type int")
	case string, bool:
		fmt.Println("x is of type string or bool")
	default:
		fmt.Printf("x is of %T\n", v)
	}
}

func learnInterface() {
	stringType := "Shubham"
	intType := 10
	userType := User{
		FirstName: "Shubham",
		LastName:  "Vyas",
	}
	// As the printInterface function takes an empty interface, We can assign any value to it
	printInterface(stringType)
	printInterface(intType)
	printInterface(userType)

	// Stringer interface
	fmt.Println(userType)

	learnTypeAssertion2()
}
