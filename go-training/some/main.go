package main

import (
	"fmt"
)

func main() {

	num := 10

	for i := 1; i <= num; i++ {
		for j := 1; j <= 10; j++ {
			res := i * j
			fmt.Println(i, "*", j, "=", res)
		}
		fmt.Println()

	}
}
