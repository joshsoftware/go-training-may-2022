package main

import "fmt"

func main() {

	//here length is defined
	var arr1 = [3]int{1, 3, 5}
	arr2 := [4]int{2, 4, 6, 8}

	fmt.Println(arr1, arr2)
	fmt.Printf("Length of an arr1 : %d\n", len(arr1))

	// here length is inferred
	var arr3 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	arr4 := [...]int{1, 3, 5, 7, 9}

	fmt.Println(arr3, arr4)
	fmt.Printf("Length of an arr3 : %d, arr4 : %d\n", len(arr3), len(arr4))

	//default array:
	var defaultArr = [4]int{}
	fmt.Println(defaultArr) //output : [0,0,0,0]

	//multi-d array:
	//default multi:d
	nums := [3][4]int{}
	fmt.Println(nums) //3 blocks with 4 elements in each block

	//string type array:
	bikes := [4]string{"Apache", "Hero", "Enfield", "Pulsor"}
	fmt.Println(bikes)

	//value type:

	x := [9]int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	copyX := x
	fmt.Println(x)
	fmt.Println(copyX)

	copyX[0] = 0
	fmt.Println(x)
	fmt.Println(copyX)
	fmt.Println(x)

	// x[0] = 0
	// fmt.Println(x)
	// fmt.Println(copyX)
}
