package main

import (
	"fmt"
)

func arrayExample() {
	array1 := [3]int{20, 30, 50}
	fmt.Println("originally, array1 is ", array1)

	// change third element , 50 to 40
	array1[2] = 40
	fmt.Println("After updation, array1 is ", array1)
	fmt.Println("Array 1 is ", array1, " with length ", len(array1), "and capacity ", cap(array1))
	array2 := array1
	fmt.Println("Array 2 is ", array2, " with length ", len(array2), "and capacity ", cap(array2))
	array3 := array2[0:2]
	fmt.Println("Array 3 is ", array3, " with length ", len(array3), "and capacity ", cap(array3))

	fmt.Println("Check if Arrays 1 and 2 are equal: ", array1 == array2)
	// make an array using make method, with length 2 smaller than capacity
	arrayMake := make([]int, 3, 5)
	fmt.Println("Array made is ", arrayMake, " with length ", len(arrayMake), "and capacity ", cap(arrayMake))

}

func sliceExample() {
	array1 := []int{}
	fmt.Println("Slice Array 1 is ", array1, " with length ", len(array1), "and capacity ", cap(array1))
	array1 = append(array1, 12)
	fmt.Println("In the slice, after append operation, Array 1 is ", array1, " with length ", len(array1), "and capacity ", cap(array1))
	array2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("In the slice, after append operation, Array 2 is ", array2, " with length ", len(array2), "and capacity ", cap(array2))
	array3 := array2[:5]
	fmt.Println("In the slice, after append operation, Array 3 is ", array3, " with length ", len(array3), "and capacity ", cap(array3))

}

func main() {
	arrayExample()
	fmt.Println()
	sliceExample()
}
