package main

import "fmt"

//Structs in Go is typed collecation of fields

//Create a struct
type person struct {
	//add some fields in structs
	firstName string
	lastName  string
	age       int
}

func main() {

	//create a simple object of struct
	var p person
	//print it
	fmt.Println(p) //print all default values of fields in struct

	//assign some value in struct and print
	person1 := person{"Rohit", "Sangamnerkar", 23}
	fmt.Println(person1)
}
