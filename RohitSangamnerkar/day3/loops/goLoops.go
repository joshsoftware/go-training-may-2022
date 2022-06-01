package main

import (
	"fmt"
)

func main() {

	//there are diffrent types of loops in go'\

	//simple for loop
	fmt.Println("Testing simple for loop and printing table of 10")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v \n", 10, i, 10*i)
	}

	//for loop as a while
	fmt.Println("Testing for as while")
	i := 5
	for i > 0 {
		fmt.Println(i)
		//decerement of i by 1
		i--
	}

	//for loop as infinite loop
	//just do not provide any condition
	/*
		for{
			//these statements will get executed infinitly
		}
	*/

	//last one is using range
	fmt.Println("Test range")
	str := "Hello"
	for _, chr := range str {
		fmt.Println(string(chr))
	}

}
