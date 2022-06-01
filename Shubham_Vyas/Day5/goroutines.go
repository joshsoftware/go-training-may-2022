package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroup() {
	// Wait group

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(s int) {
			fmt.Println(s)
			wg.Done()
		}(i)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("Hello")
}

// Now expected output is 1000 but we are getting 876
// This output would vary for every run as goroutines runs at random
// The reason is some routines are using same value when intialized
// To prevent such race condition we use locks
// locks can be implemented using mutex or channels
func learnGoRoutine() {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			x += 1
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Value of x is ", x)
}
