package main

import "fmt"

func main() {
	pi, euler := 3.142, 2.718
	pointer := &pi
	fmt.Println("value of pi by deferencing", *pointer)
	fmt.Println("value of pointer is", pointer)

	pointer = &euler
	fmt.Println("p is now pointing to euler", *pointer)

	fmt.Println("value of pointer is", pointer)
}
