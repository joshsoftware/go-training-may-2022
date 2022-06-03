package main

import (
	"fmt"
	"time"
)

//wg sync.WaitGroup
func raceConditionExample() {
	sum := 0
	for i := 0; i < 1000; i++ {
		go func() {
			sum = sum + 1
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("sum is -->", sum)
}

func main() {
	raceConditionExample()
}
