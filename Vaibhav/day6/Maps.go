package main

import "fmt"

func main() {

	// Creating a map
	// map[key-type]value-type
	// var Yashwin = make(map[int]string)
	// fmt.Println(Yashwin)

	// // Adding elements to the map
	// Yashwin[1] = "Rohit"
	// Yashwin[2] = "Sumit"
	// Yashwin[3] = "Vaibhav"
	// Yashwin[4] = "Rajat"

	// fmt.Println(Yashwin)

	countries := map[string]string{
		"IN": "India",
		"NP": "Nepal",
		"TR": "Turkey",
		"JP": "Japan",
		"ZW": "Zimbabwe",
	}

	for country := range countries {
		fmt.Println(country, "=>", countries[country])
	}

	for key, value := range countries {
		fmt.Printf("countries[%s] = %s\n", key, value)
	}

}
