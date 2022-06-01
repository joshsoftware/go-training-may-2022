package main

import "fmt"

func main() {
	fmt.Println("Hello From go")
	i, j := 20, 15
	p := &i
	fmt.Println("p is pointing to i", *p)

	p = &j
	fmt.Println("p is pointing to j", *p)
	*p = 32
	fmt.Println("value of j is", j)
	fmt.Println("p is pointing to j", *p)
}
