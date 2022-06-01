package main

import "fmt"

func main() {
	//In go arrays are pass by value
	//we can declare array in two ways
	//first
	var arr1 [3]int //arr1 variable is array of size 3 of integer
	//second short hand operator and initialize at same time
	arr2 := [3]int{10, 20, 30}
	//print secont array
	fmt.Println("second array : ", arr2)
	//initialize first array
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30

	//print first array
	fmt.Println("First Array : ", arr1)

	//arrays are value type so we can compare two arrays using
	fmt.Println(arr1 == arr2)

}
