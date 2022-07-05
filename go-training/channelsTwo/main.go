package main

import (
	"fmt"
	"sync"
)

func main() {

	//
	wg := &sync.WaitGroup{}
	ch := make(chan int) //channel

	wg.Add(2)
	// receiver
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(ch)
		val, isChannelOpen := <-ch
		fmt.Println(val)
		fmt.Println(isChannelOpen)
		wg.Done()
	}(ch, wg)

	//sender
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 8
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait() //blocker
}
