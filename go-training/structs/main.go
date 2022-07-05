package main

import "fmt"

//struct
type Person struct {
	name string
	age  int
	job  string
}

//nested struct
type person struct {
	name string
	age  int
}
type engineer struct {
	details person
}

func main() {

	var pers1 Person
	var pers2 Person

	//pers1
	pers1.name = "Rak"
	pers1.age = 26
	pers1.job = "Software Engineer"

	//pers2
	pers2.name = "Suresh"
	pers2.age = 29
	pers2.job = "Bank Employee"

	fmt.Println(pers1)
	fmt.Println(pers2)

	//nested struct
	res := engineer{
		details: person{"Raj", 26},
	}
	fmt.Println(res.details)
}
