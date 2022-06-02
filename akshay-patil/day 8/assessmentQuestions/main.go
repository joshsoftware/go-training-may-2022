package main

import "fmt"

func main() {
	//switchCaseExample()
	SliceExample()
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

//slices are tricky
func SliceExample() {
	colors := []string{"red", "blue"}
	addingElement(colors)
	fmt.Println("after addingElement --> ", colors)
	addingElementUsingPointer(&colors)
	fmt.Println("after addingElementUsingPointer --> ", colors)
	changingElement(colors)
	fmt.Println("after changingElement--> ", colors)
}

//
func addingElement(colors []string) {
	colors = append(colors, "green")
	fmt.Println("colors--> ", colors)
}

func addingElementUsingPointer(colors *[]string) {
	*colors = append(*colors, "green")
	fmt.Println("colors--> ", *colors)
}

func changingElement(colors []string) {
	colors[0] = "yellow"
	fmt.Println("colors--> ", colors)
}
