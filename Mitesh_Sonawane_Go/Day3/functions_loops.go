package main

import "fmt"

func Swap(x, y string) (string, string) {
	return y, x
}

func Max(num1, num2 int) int {
	var result int
	if num1 != num2 {
		if num1 > num2 {
			result = num1
		} else {
			result = num2
		}
	} else {
		fmt.Println("num1 and num2 are same")
	}
	return result
}

func main() {
	a := 5
	// method I for if else
	if a > 10 {
		fmt.Printf("%v is greater than 10 \n", a)

	} else {
		fmt.Printf("%v is less than 10 \n", a)
	}

	// method II for if else
	if a > 10 {
		fmt.Printf("%v is greater than 10 \n", a)
		return
	}
	fmt.Printf("%v is less than 10 \n", a)

	fmt.Println("****************")

	// name, sirname := "Mayur", "Sonawane"
	// fmt.Println(Swap(name, sirname))

	name, sirname := Swap("Mayur", "Sonawane")
	fmt.Println(name, sirname)

	fmt.Println(Swap("mitesh", "Sonawane"))
	fmt.Println(Max(11, 55))

	fmt.Println("***  for Loop  ****")

	for i := 0; i <= 10; i++ {
		fmt.Printf("%v \t", i)
	}
	fmt.Println()

	range_upto := []int{11, 22, 33, 44, 55}
	for i, num := range range_upto {
		fmt.Println(i, num)
	}

}
