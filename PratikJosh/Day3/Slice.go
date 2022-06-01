package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := []int{2, 3, 5, 7, 11, 13}

	var a []int = arr[1:4]
	fmt.Println(a)

	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
