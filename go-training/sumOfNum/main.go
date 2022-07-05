// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	num := 671
	n := 0
	for n != 0 {
		n = n + num%10
		println(n)
		num = num / 10
		println(num)

	}
	fmt.Println(n)
}
