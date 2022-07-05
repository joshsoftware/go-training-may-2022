package main

import "fmt"

func main() {

	a := 8
	var ptr1 *int
	ptr1 = &a
	fmt.Println(`Address of "a" :`, &a)
	fmt.Println(`Address of "a through Pointer :`, ptr1)
	fmt.Println(`Value of "a" throug Pointer :`, *ptr1)

	//short hand

	b := 6
	ptr2 := &b
	fmt.Println(`Address of "b" :`, &b)
	fmt.Println(`Address of "b" through Pointer :`, ptr2)
	fmt.Println(`Value of "b" throug Pointer :`, *ptr2)

}
