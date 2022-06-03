package main

import "fmt"

func main() {
	//switchCaseExample()
	sliceExample()
	//namedReturnExample()
}

func switchCaseExample() {
	// In go, even if we don't give the expression in switch case it will still run
	// the first true statement
	switch {
	case true:
		fmt.Println("it's true")
	case 1 < 3:
		fmt.Println("it's 1<3")
	case false:
		fmt.Println("it's false")
	default:
		fmt.Println("it's default")
	}
}

// slices are tricky
func sliceExample() {
	colors := []string{"red", "blue"}
	addingElement(colors)
	fmt.Println("after addingElement --> ", colors)
	addingElementUsingPointer(&colors)
	fmt.Println("after addingElementUsingPointer --> ", colors)
	changingElement(colors)
	fmt.Println("after changingElement--> ", colors)
}

// In the addingElement() method, if we try to append the slice,
// it seems to be creating a new slice which is of local scope
func addingElement(colors []string) {
	colors = append(colors, "green")
	fmt.Println("colors--> ", colors)
}

// In the addingElementUsingPointer() method, if we try to append
// the slice, it will do the append to the actual slice too!
func addingElementUsingPointer(colors *[]string) {
	*colors = append(*colors, "green")
	fmt.Println("colors--> ", *colors)
}

// In the changingElement() method, if we try to change the elements
// from the slice, it will change the original slice elements too!!
func changingElement(colors []string) {
	colors[0] = "yellow"
	fmt.Println("colors--> ", colors)
}

func namedReturnExample() {
	s := sum(10, 12)
	fmt.Println("sum is --> ", s)
}

// In name return type, if we assign the result value to the return var name,
// we don't even have to give anything after return statement.
// And if we give any other var after return statement, then it should be same type.
func sum(a, b int) (sum *int) {
	s := a + b
	return &s
}

// If we only have a foo() and if go package/file doesn't contains the main(), the
// program will not execute and will get Build Failure. Because main() is an entry point of the program.
// But in failed scenario, program execution will panic??
func foo() {}
