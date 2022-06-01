package main

import "fmt"

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	s := arr[3:6]
	fmt.Println(s)

	s = append(s, 11)
	s = append(s, 12)
	s = append(s, 13)

	fmt.Println(arr)
	fmt.Println(s)

	fmt.Println("*******************")

	arr1 := []int{1, 2, 3, 4, 5}
	arr1 = append(arr1, 55)
	fmt.Println(arr1)

	s1 := arr1[1:4]
	fmt.Println(s1)

	s1 = append(s1, []int{66, 77, 88}...)
	fmt.Println(s1)
	fmt.Println(arr1)

	fmt.Println("***********")
	a1 := [3]int{1, 4, 3}
	a2 := [3]int{1, 2, 3}
	fmt.Println(a1 == a2)

}
