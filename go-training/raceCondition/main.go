package main

import (
	"fmt"
	"sync"
)

func main() {

	//Race Condition:
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	//creating one slice
	var nums = []int{0}

	//goroutines
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One")
		mut.Lock()
		nums = append(nums, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two")
		mut.Lock()
		nums = append(nums, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three")
		mut.Lock()
		nums = append(nums, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(nums)
	fmt.Println("Done!!")
}
