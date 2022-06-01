package main

import "fmt"

var arr = [8]int{1, 2, 3, 4, 5, 6, 7, 8}

// This example shows that when we create a slice from an array, it actually peeks into array.
func sliceExample1() {
	// Here we are taking all elements from index 3 to index 5.
	// Remember slice wont take 6th index.
	slice := arr[3:6]

	// If I change anything in slice it will be reflected in array.
	// As slice will be created with a ref type

	slice[2] = 10

	fmt.Println("Slice: ", slice)
	fmt.Println("Array: ", arr)

	fmt.Printf("Length of slice: %d & Capacity of slice: %d", len(slice), cap(slice))
}

// Lets append a new element to slice
func sliceExample2() {
	slice := arr[3:6]
	slice = append(slice, 11)

	// In result you will c that slice and array both will be changed
	// New elem will be added in slice
	// And value of element at position 6 in array will be changed
	fmt.Println("Slice: ", slice)
	fmt.Println("Array: ", arr)
}

func sliceExample3() {
	slice := arr[3:6]
	// The ... is ellipsis operator which will add all elements to slice instead of adding slice
	// kind of spread operator in js
	slice = append(slice, []int{11, 12, 13, 14}...)

	// Here array wont change but slice will.
	// As slice is exceeding capacity of array, slice ref will be changed from arr to a new slice with updated capacity
	// That the reason old arr wont be changed anymore
	// Remember a slice has [len, cap, pointer], pointer will hold ref for the data
	fmt.Println("Slice: ", slice)
	fmt.Println("Array: ", arr)
}
