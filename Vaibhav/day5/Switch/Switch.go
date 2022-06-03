package main

import "fmt"

func main() {

	switch {
	case true:
		fmt.Println("in true case")
	case 1 < 2:
		fmt.Println("in 1<2 case")
	case false:
		fmt.Println("in false case")
	default:
		fmt.Println("in default case")
	}

}
