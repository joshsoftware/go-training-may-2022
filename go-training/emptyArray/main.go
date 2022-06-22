package main

import "fmt"

func main() {

	//declaring an empty array:
	var arrayOfInt [4]int
	fmt.Printf("Type : %T, Values :%d\n", arrayOfInt, arrayOfInt)

	//assigning elements into an array using index
	arrayOfInt[0] = 2
	arrayOfInt[1] = 4
	arrayOfInt[2] = 6
	arrayOfInt[3] = 8

	fmt.Println(arrayOfInt)

	//changing values
	arrayOfInt[0] = 9
	fmt.Println(arrayOfInt)

	//length of an array:
	lenOfArray := len(arrayOfInt)
	fmt.Println("Length of an Array is ", lenOfArray)

	//looping over an array:
	for index, value := range arrayOfInt {
		fmt.Println(index, ":", value)
	}
	fmt.Println()

	//looping with only index
	for i := range arrayOfInt {
		fmt.Println(arrayOfInt[i], ":", i)
	}

	//accesing only values
	for a := 0; a < lenOfArray; a++ {
		fmt.Println("Values are :", arrayOfInt[a])
	}

	//multi-dimentional array:
	multiArray := [2][3]int{{1, 3, 5}, {2, 4, 6}}
	fmt.Println(multiArray)

	for x := 0; x < len(multiArray); x++ {
		for y := 0; y < len(multiArray); y++ {
			fmt.Println(multiArray[x][y])
		}
	}
}
