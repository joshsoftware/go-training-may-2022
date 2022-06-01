package main

import "fmt"

func main() {
	a := 5
		// method I for if else
	if a > 10 {
		fmt.Printf("%v is greater than 10 \n", a)

	} else {
		fmt.Printf("%v is less than 10 \n", a)
	}

		// method II for if else
	if a > 10 {
		fmt.Printf("%v is greater than 10 \n", a)
		return
	}
	fmt.Printf("%v is less than 10 \n", a)

}
