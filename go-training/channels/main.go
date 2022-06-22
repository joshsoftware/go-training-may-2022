package main

import "fmt"

func main() {

	//creating channels :
	name := make(chan string)
	age := make(chan int)

	//function call with goroutine
	go sendData(name, age)

	//received channel data
	Name := <-name
	Age := <-age
	fmt.Println("Name :", Name, "\nAge :", Age)
}

func sendData(name chan string, age chan int) {

	//send data to the channels:
	name <- "Rajkumar Bingi"
	age <- 26

	fmt.Println("No Receiver at the moment")
}
