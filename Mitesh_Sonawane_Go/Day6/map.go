package main

import "fmt"

func main() {
	var mp = map[string]int{
		"mitesh":  100,
		"mayur":   200,
		"bhushan": 50,
	}

	// to see value present or not if present then true
	// else false and gives value of key
	val, ok := mp["mitesh"]
	fmt.Println(val, ok)

	mp["mayur"] = 333
	fmt.Println(mp)

	for key, values := range mp {
		// _ is blank identifier
		fmt.Println(key, values)
	}

	fmt.Print(" total keys = ")
	for key, _ := range mp {
		// _ is blank identifier
		fmt.Print(" ", key)
	}

	fmt.Println("\n****  method II *********")
	var salary = make(map[string]float64)
	salary["jay"] = 2233.5
	salary["raj"] = 555.6666

	fmt.Println(mp)
	fmt.Println(salary)
}
