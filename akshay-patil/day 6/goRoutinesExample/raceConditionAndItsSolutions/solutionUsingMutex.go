package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	sum := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mx.Lock()
			sum = sum + 1
			mx.Unlock()
			wg.Done()
		}()
		wg.Add(1)
	}
	time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("sum is -->", sum)

}
