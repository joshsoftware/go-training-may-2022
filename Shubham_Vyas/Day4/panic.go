package main

import "fmt"

// panic(interface{}) it takes any concrete type as argument
// panic will abort current process
// Mostly should be avoided, unless necessary

type User struct {
	FirstName string
	LastName  string
}

func panicInt() {
	i := 10
	panic(i)
	// Output - panic: 10
}

func panicString() {
	s := "Check whats failing"
	panic(s)
	// Output - panic: Check whats failing
}

func panicUser() {
	user := User{
		FirstName: "Shubham",
		LastName:  "Vyas",
	}
	panic(user)
	// panic: (main.User) 0xc000068020
	// above will be output
}

// defer will execute even when panic occurs
func deferWithPanic() {
	defer fmt.Println("Hello before panic")
	panic("Panic")
	// Output - Hello before panic\npanic: Panic
}
func leanPanic() {
	// panicUser()
	// panicString()
	// panicInt()
	deferWithPanic()
}
