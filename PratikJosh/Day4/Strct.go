package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println("we are learning structs  here")
	pratik := User{"pratik", 23, "pratikvarute07@gmail.com"}
	fmt.Println(pratik)
	fmt.Println("hey there pratik details are:  %+v\n", pratik)
}
