package main

import "fmt"

func main() {
	//switchCaseExample()
	//sliceExample()
	namedReturnExample()
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
// it seems to be creating the new slice which seems to be having
// method level scope
func addingElement(colors []string) {
	colors = append(colors, "green")
	fmt.Println("colors--> ", colors)
}

// In the addingElementUsingPointer() method, if we try to append
// the slice, it will do the append to the actual array too!
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
// And if we give any other name after return statement, then it should be same type.
func sum(a, b int) (sum *int) {
	s := a + b
	return &s
}
