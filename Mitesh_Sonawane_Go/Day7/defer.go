package main

import "fmt"

func main() {
	/*
			fmt.Println("hello ")

			for i := 1; i <= 3; i++ {
				defer fmt.Println(i)
			}


		fmt.Println("world")
	*/

	/*
		x := 11
		defer fmt.Println(x)
		x = 22

		fmt.Println(x)

	*/

	a := 4
	b := 6

	// defer func() {
	// 	fmt.Println(a + b)
	// }()

	b = 5
	defer fmt.Println(a + b)
	return

}
