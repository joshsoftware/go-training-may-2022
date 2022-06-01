package main

import "fmt"

func compare() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	// Answer will be true as we can compare array data type.
	fmt.Println(arr1 == arr2)

	// slice1 := []int{1, 2, 3}
	// slice2 := []int{1, 2, 3}
	// fmt.Println(slice1 == slice2)

	// Slices cannot be compared in go
	// Slice, Map and Channel stuct cannot be compared in go
	// Uncomment above snippet and check error

	// Array can be compared only when there type matches
	// Type of array is [3]int = array of int of size 3

	// uncomment
	// arr3 := [2]int{1, 2}
	// fmt.Println(arr1 == arr3)
}
