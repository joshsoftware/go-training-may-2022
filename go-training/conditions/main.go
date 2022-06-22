package main

import "fmt"

func main() {

	//if : print even numbers
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		}
	}

	//if-else : find even or odd
	num := 7
	if num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}

	// if-else-if :
	n := 8
	if n < 0 {
		fmt.Println(n, " is negitive")
	} else if n == 0 {
		fmt.Println(n, " is 0")
	} else {
		fmt.Println(n, " is positive")
	}

	// nested if :
	a := 0
	if a >= 0 {
		if a == 0 {
			fmt.Println(a, " is zero")
		} else if a < 0 {
			fmt.Println(a, " is negitive")
		} else {
			fmt.Println(a, " is postive")
		}
	}

	// switch case :
	day := 8
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednessday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Its not a week day!!")

	}

}
