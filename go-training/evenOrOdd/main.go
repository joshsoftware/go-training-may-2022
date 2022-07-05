package main

import "fmt"

func main() {

	var num int

	fmt.Print("Enter Number : ")
	fmt.Scanln(&num)

	if num > 0 {
		if num%2 == 0 {
			fmt.Printf("%d is Even\n", num)
		} else {
			fmt.Printf("%d is Odd\n", num)
		}
	} else {
		fmt.Println("It is Zero")
	}
}
