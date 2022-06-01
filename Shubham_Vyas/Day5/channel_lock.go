package main

import (
	"fmt"
	"sync"
)

func incrementUsingChannel(i int, ch chan int, wg *sync.WaitGroup) {
	// unless der is a reader we wont right in channel
	// So it will block
	ch <- i + 1
	wg.Done()
}

func learnBufferedChannel() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	val := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementUsingChannel(val, ch, &wg)
		// unless der is writer we wont read from channel
		val = <-ch
	}

	wg.Wait()
	fmt.Println(val)
}
