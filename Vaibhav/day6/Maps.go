package main

import "fmt"

func main() {

	// Creating a map
	// map[key-type]value-type
	var Yashwin = make(map[int]string)
	fmt.Println(Yashwin)

	// Adding elements to the map
	Yashwin[1] = "Rohit"
	Yashwin[2] = "Sumit"
	Yashwin[3] = "Vaibhav"
	Yashwin[4] = "Rajat"

	fmt.Println(Yashwin)

}
