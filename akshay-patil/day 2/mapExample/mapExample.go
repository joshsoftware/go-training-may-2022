package main

import "fmt"

func main() {
	//map using make keyword
	makeMap := make(map[int]int, 3)
	fmt.Println("mapUsingMake -->", len(makeMap))
}
