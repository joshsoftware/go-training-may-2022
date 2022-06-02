package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}
type Emp struct {
	Name  string
	EID   int
	Email string
}

func main() {
	fmt.Println("we are learning structs  here")
	pratik := User{"pratik", 23, "pratikvarute07@gmail.com"}
	pratham := Emp{"pratham", 23, "prathm@gmail.com"}
	fmt.Println(pratik)
	fmt.Println(pratham)
	fmt.Println("hey there pratik details are:  %+v\n", pratik)
}
