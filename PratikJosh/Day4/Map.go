package main

import (
	"fmt"
)

func main() {

	languges := make(map[string]string)

	languges["js"] = "javascript"
	languges["rb"] = "ruby"
	languges["go"] = "Golang"

	//fmt.Println(languges["rb"])

	//delete(languges, "rb")

	fmt.Println(languges["rb"])
	for key, value := range languges {
		fmt.Println("for %v key is %v", key, value)

	}

}
