package main

import "fmt"

func main() {
	//allocating memory and initializing map using make function
	m := make(map[string]int)
	//if you want to add values then you need
	m["Rohit"] = 45000
	m["Vaibhav"] = 55000
	m["Rajat"] = 45000
	m["Sumit"] = 45000
	fmt.Println(m)
	//second way to init map is
	m1 := map[string]int{
		"Rohit":   1,
		"Sumit":   2,
		"Vaibhav": 3,
		"Rajat":   4,
	}

	//print new map
	fmt.Println(m1)
}
