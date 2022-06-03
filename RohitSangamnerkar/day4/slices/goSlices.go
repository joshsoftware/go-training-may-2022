package main

import "fmt"

func main() {

	//we can also make slice using make function
	//if we dont give size here then is acts as slice
	//and it  is dynamic in nature
	slice1 := []int{10, 20}

	//we can also create slice from another array

	//lets declare on array
	arr := [5]int{10, 20, 30, 40, 50}
	//to make a slice from above array
	slice := arr[1:3]
	//here sice of slice is 2
	//And capacity of slice is 4

	//now just print array and slice
	fmt.Println("Array : ", arr)
	fmt.Println("Slice : ", slice)
	//

	//we cannot compare two slices by using == operator

	fmt.Println(slice1)
}
