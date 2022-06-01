package main

import (
	"fmt"
	"time"
)

func doSomething1(s string, ch chan int) {
	ch <- 8
	fmt.Println(s)
}
func doSomething2(s string, ch chan int) {
	i := <-ch
	fmt.Println(s)
	fmt.Println(i)
}
func main() {

	ch := make(chan int)
	go doSomething1("Hello", ch)
	go doSomething2("World", ch)

	// Unbuffered channels are always passed by reference
	// ch := make(chan int)
	// ch <- 10
	// i := <-ch
	// fmt.Println(i)
	time.Sleep(time.Second * 2)
}
