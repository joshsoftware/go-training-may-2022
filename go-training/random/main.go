package main

import (
	// "math/rand"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	//random number from math:
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(6) + 1)

	//random number from crypto
	myRandomNumber, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		fmt.Println("Oops!!")
	}
	fmt.Println(myRandomNumber)
}
