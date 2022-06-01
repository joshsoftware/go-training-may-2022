package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func GoRoutines() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Go routine completed")
}

func Channels() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)
	go sum(s, c)
	result := <-c
	fmt.Println("result throw channel ", result)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum = sum + v
	}
	c <- sum
}

func main() {
	GoRoutines()
	Channels()
}
