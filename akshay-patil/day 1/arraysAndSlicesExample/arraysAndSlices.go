package main

import "fmt"

func ArrayExample() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1 is --> ", arr1)
	arr1[3] = 40
	fmt.Println("arr1 after chaning  is ----> ", arr1)
	arr2 := arr1
	fmt.Println("arr2 --> ", arr2, " length is --> ", len(arr2), "capacity is --> ", cap(arr2))
	arr3 := arr2[1:3]
	fmt.Println("arr3 --> ", arr3, " length is --> ", len(arr3), "capacity is --> ", cap(arr3))
	//below line will give panic: runtime error: index out of range [4] with length 2
	//arr3[2] = 32
	arr4 := arr1
	fmt.Println(arr4)
}

func SliceExample() {
	arr1 := []int{}
	fmt.Println("In slice arr1 --> ", arr1, " length is --> ", len(arr1), "capacity is --> ", cap(arr1))
	arr1 = append(arr1, 12)
	fmt.Println("In slice, After append arr1 --> ", arr1, " length is --> ", len(arr1), "capacity is --> ", cap(arr1))
	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("In slice, After append arr2 --> ", arr2, " length is --> ", len(arr2), "capacity is --> ", cap(arr2))
	arr3 := arr2[:5]
	fmt.Println("In slice, After append arr3 --> ", arr3, " length is --> ", len(arr3), "capacity is --> ", cap(arr3))
	//fmt.Println(arr3[:cap(arr3)])

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	//it can compare the array, but in case of slice it is not comparing
	fmt.Println(a1 == a2)
}

func main() {
	ArrayExample()
	SliceExample()
}
