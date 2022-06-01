package main

import "fmt"

// Output would be Hello than World
func learnDefer1() {
	// defer will delay the execution untill function returns
	defer fmt.Println("World")
	fmt.Println("Hello")
}

// Output will be 20
func learnDefer2() {
	a, b := 10, 4
	defer func() {
		fmt.Println(a + b)
	}()
	b = 10
}

// output will be 24
func learnDefer3() {
	a, b := 10, 4
	// c cannot be accessed outside function
	defer func() {
		c := a + b
		fmt.Println(c)
	}()
	b += 10
}

// output will be 14
func learnDefer4() {
	a, b := 10, 4

	// Here this function is not using a,b directly
	// It will hold value of c till it executes as while evaluating we are passing 14 as a + b
	defer func(c int) {
		fmt.Println(c)
	}(a + b)

	b += 10
}

// Output will be 9,8,7...
func learnDefer5() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
