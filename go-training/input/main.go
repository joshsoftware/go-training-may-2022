package main

import "fmt"

func main() {

	// Scan():
	// var num int
	// fmt.Print("Enter a number : ")
	// fmt.Scan(&num)

	// fmt.Println("Number is ", num)

	//Scanln :
	// var name string
	// var age int

	// fmt.Println("Enter Name and Age :")
	// fmt.Scanln(&name, &age)

	// fmt.Printf("Name : %s,Age : %d\n", name, age)

	//Scanf
	var name string
	var age int

	fmt.Println("Enter Name and Age :")
	fmt.Scanf("%s,%d", &name, &age)

	fmt.Printf("Name : %s,Age : %d\n", name, age)

}
