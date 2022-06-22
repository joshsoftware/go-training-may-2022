package main

import "fmt"

func main() {

	var a = map[string]string{"name": "Rak", "Age": "26"}
	b := map[int]string{1: "One", 2: "Two", 3: "Three"}

	fmt.Println(a)
	fmt.Println(b)

	//create a map using make():
	var cars = make(map[string]string) //empty map
	fmt.Println(cars)

	cars["Brand"] = "Jeep"
	cars["model"] = "Meridian"

	fmt.Println(cars)

	// map for few actions
	names := map[int]string{1: "Rak", 2: "Reddy", 3: "Sahu", 4: "Turbo", 5: "Ajay"}
	fmt.Println(names)

	//adding
	names[6] = "xyz"
	fmt.Println(names)

	//update
	names[6] = "abcd"
	fmt.Println(names)

	delete(names, 6)
	fmt.Println(names)

	//access through key :
	res := names[1]
	fmt.Println(res)

	//checking : val = value ,ok = true/false
	val1, ok1 := names[1]
	val2, ok2 := names[8]
	_, ok3 := names[2]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(ok3)

	//maps are reference types:
	copyNames := names
	fmt.Println(copyNames)

	copyNames[6] = "Madhu"
	fmt.Println(copyNames)
	fmt.Println(names)

	// looping over maps:
	for key, value := range names {
		fmt.Println(key, ":", value)
	}

}
