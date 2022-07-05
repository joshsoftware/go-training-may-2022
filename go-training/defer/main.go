package main

import "fmt"

func main() {

	// for i := 1; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%d is even number\n", i)
	// 	} else {
	// 		defer fmt.Printf("%d is odd number\n", i)
	// 	}
	// }

	// fmt.Println("start of function..")
	// for i := 0; i < 8; i++ {
	// 	defer fmt.Println("Good bye..!!")
	// 	fmt.Println(1)
	// 	fmt.Println(2)
	// 	fmt.Println(3)
	// 	fmt.Println(4)
	// 	return
	// 	fmt.Println(5)
	// 	fmt.Println(6)
	// 	fmt.Println(7)
	// 	fmt.Println(8)
	// }

	fmt.Println("Start")

	i := 23
	defer fmt.Println(`value of "i" :`, i)

	i = 33

	fmt.Println(foo())
}

func foo() (result string) {
	defer func() {
		result = "Change World"
	}()
	return "Helloworld"
}
