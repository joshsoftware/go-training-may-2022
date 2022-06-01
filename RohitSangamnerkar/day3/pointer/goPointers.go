package main

import "fmt"

func main() {
	//pointers in golang is special variable
	//which holds a address of another veriable

	//declare and initialize a veriable
	a := 10

	var pointer *int = &a

	fmt.Println(a)        //prints value in a
	fmt.Println(&a)       //prints address of a
	fmt.Println(pointer)  //this prints value of pointer which is addres of a
	fmt.Println(*pointer) //it will print value at address which pointer holds

}
