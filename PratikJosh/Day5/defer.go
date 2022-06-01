package main

import "fmt"

func main() {
	i := 23
	defer fmt.Println("I=", i)
	i = 33
	defer fmt.Println("I=", i)
}
