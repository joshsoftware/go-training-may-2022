package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// func Addd(x, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }

func foo(x int) {
	for i := 0; i <= x; i++ {
		fmt.Println(i)
	}
}

func bar() {
	for i := 600; i <= 1000; i++ {
		fmt.Println(i)
	}
}

func main() {

	//time.Sleep(2 + time.Second)
	wg.Wait()
	go foo(500)
	bar()

	//Addd(11, 11)
}
