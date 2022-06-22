package main

import "fmt"

func main() {
	res := 0

	for i := 0; i < 10; i++ {
		res += i
	}
	fmt.Println(res)

	// while in for
	i := 0
	for i < 5 {
		i++
	}
	fmt.Println(i)

	//for each
	names := []string{"Rak", "Sahu", "Reddy", "Kachi", "Turbo"}
	for i := range names {
		fmt.Println(names[i])
	}
}
