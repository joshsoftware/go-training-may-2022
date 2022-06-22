package main

import "fmt"

func main() {

	//int type slice
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	//string type slice
	cars := []string{"Benz", "Audi", "Jeep", "MG"}
	fmt.Println(cars)

	//create an slice from an array :
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersSlice := numbers[2:8] //index value : length of an array
	copyNumbers := numbers[:]
	fmt.Println(numbersSlice)
	fmt.Println(copyNumbers)

	//create a slice using make()
	a := make([]int, 3, 4) // 3: length , 4: capacity
	fmt.Println(a)
	a[0] = 10
	a[1] = 20
	a[2] = 30
	// a[3] = 40
	fmt.Println(a)
	fmt.Printf("length : %d,capacity : %d\n", len(a), cap(a))

	// append():
	a = append(a, 40, 50)
	fmt.Println(a)
	fmt.Printf("length : %d,capacity : %d\n", len(a), cap(a))

	// appending slice into another slice:
	b := []int{60, 70, 80}
	b = append(a, b...)
	fmt.Println("slice append :", b)
	fmt.Printf("length : %d,capacity : %d\n", len(b), cap(b))

	// copy():
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ele := n[:6]
	new := make([]int, len(ele))
	copy(new, ele)
	fmt.Println(new)

	//ref type:

	x := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	copyX := x[:5]
	fmt.Println(x)
	fmt.Println(copyX)

	copyX[0] = 0
	fmt.Println(x)
	fmt.Println(copyX)

	x[0] = 2
	fmt.Println(x)
	fmt.Println(copyX)
}
