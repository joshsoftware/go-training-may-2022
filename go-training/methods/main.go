package main

import "fmt"

type rect struct {
	height, width int
}

func (r rect) area() int {
	return r.height * r.width
}
func (r *rect) perim() int {
	return 2 * r.height * r.width
}

func main() {
	r1 := rect{height: 6, width: 8}
	fmt.Println("Area :", r1.area())
	fmt.Println("Perimeter :", r1.perim())

	//pointer
	r2 := &r1
	fmt.Println("Area :", r2.area())
	fmt.Println("Perimeter :", r2.perim())

}
