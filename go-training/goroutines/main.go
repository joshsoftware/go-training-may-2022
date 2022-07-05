package main

import (
	"fmt"
	"sync"
)

// main function without wait-gorup
// func main() {

// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Goroutines")
// }

//WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup //pointer

//main function with WaitGroup
func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Add(1)
	}

	wg.Wait() //blocker call
	fmt.Println("Hello World!")
}
