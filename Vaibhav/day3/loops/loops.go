package main

import "fmt"

func main() {

	numbers := [5]int{1, 2, 3, 4, 5}

	//for loop execution
	for i := 0; i < 5; i++ {
		fmt.Printf("value of i: %d\n", i)
	}

	//range
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}
