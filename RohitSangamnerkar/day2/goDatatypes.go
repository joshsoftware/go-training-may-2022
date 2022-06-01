package main

import "fmt"

func main() {

	//types of datatypes
	//type1 Basic : Numbers, strings, and booleans

	//we can declare veriables in two ways short hand and using var
	//Numeric

	//for integer int8,int16,int32,int64
	//for unsigned integer uint8,uint16,uint32,uint64

	var intType int64 = 10 //using var
	intType2 := 20         //short hand

	//for float : float32 and float64

	var floatType float64 = 10.2
	flotType2 := 20.2

	fmt.Println("Printitng Numeric type : ", intType, " ", intType2, " ", floatType, " ", flotType2)

	//String datatype
	var stringType string = "test string type"
	stringType2 := "using short hand string"

	//printing
	fmt.Printf("printing String type : %s , \n %s \n", stringType, stringType2)

	//boolean data type
	//it contains only ont bit
	//it holds value true or false

	var checkBoolean bool = false
	checkBoolean2 := true

	fmt.Println("Printing boolean values : ", checkBoolean, " , ", checkBoolean2)

}
