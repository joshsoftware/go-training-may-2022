package main

import "fmt"

func main() {
	//map using make keyword
	mapUsingMake := make(map[int]int, 3)
	fmt.Println("mapUsingMake -->", len(mapUsingMake))
}
