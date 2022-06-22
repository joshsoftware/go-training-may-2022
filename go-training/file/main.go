package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Files in Golang")
	content := "Go programming language."

	//creating file
	file, err := os.Create(".mygofile.txt")

	//writing content inside file.
	checkNillErr(err)
	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println(length) //length of content inside file

	readFile(".mygofile.txt")
}

//read file:
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNillErr(err)

	fmt.Println("Data as bytes :", databyte)
	fmt.Println("Actual data :", string(databyte))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
