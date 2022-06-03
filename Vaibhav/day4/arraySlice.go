package main

import "fmt"

func main() {

	//array
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	//slice
	myslice := arr2[2:4] // myslice arr{6,7}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(myslice)

}
