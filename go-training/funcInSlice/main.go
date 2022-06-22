package main

import "fmt"

func main() {

	primeNumbers := []int{2, 3, 5}

	//append():
	primeNumbers = append(primeNumbers, 7, 11)
	fmt.Println(primeNumbers)

	//copy():
	evenNumbers := []int{12, 14, 16}
	copy(primeNumbers, evenNumbers)
	fmt.Println(primeNumbers)

	//len():
	fmt.Println(len(primeNumbers))
}
