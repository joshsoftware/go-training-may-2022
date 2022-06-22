package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Microsecond)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Microsecond)
		c2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println(m1, i)
		case m2 := <-c2:
			fmt.Println(m2, i)
		}
	}
}
