package main

import (
	"fmt"
	"sync"
)

func incrementUsingChannel(i int, ch chan int, wg *sync.WaitGroup) {
	ch <- i + 1
	wg.Done()
}

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	val := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementUsingChannel(val, ch, &wg)
		val = <-ch
	}

	wg.Wait()
	fmt.Println(val)
}
