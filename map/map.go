package main

import "fmt"

func createMap(size, index int) []int {
	//make a map using map keyword
	makeMap := make(map[int]int, size)
	mapInfo := [2]int{makeMap[index], len(makeMap)}
	return mapInfo[:]
}
func main() {
	fmt.Println(createMap(3, 2))
}
