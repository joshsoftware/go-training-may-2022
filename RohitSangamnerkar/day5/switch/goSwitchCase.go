package main

import "fmt"

//A switch statement is a shorter way to write a sequence of if - else statements.
// It runs the first case whose value is equal to the condition expression.
//unlike other languages we donot need break statment in switch case
//it will check if matched then execute statment and stop
//if not then go for next case

func main() {

	switch x := 3; x {
	case 1:
		fmt.Println("first case")
	case 2:
		fmt.Println("Second case")
	case 3:
		fmt.Println("Third case")
	case 4:
		fmt.Println("Forth case")
	default:
		fmt.Println("Default case")
	}

	//there is another type
	//like switch with no condition
	//it is like switch true

	x := 20

	switch {
	case x < 10:
		fmt.Println("less than ten")
	case x > 10 && x < 20:
		fmt.Println("Less than 20")
	default:
		fmt.Println("greater than 20")
	}

}
