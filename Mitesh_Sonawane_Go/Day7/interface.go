package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  int
	Breadth int
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return float64(r.Length) * float64(r.Breadth)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Stringer Interface
func (r Rectangle) String() string {
	return fmt.Sprintf(" Rectangle length = %v and Breadth = %v ", r.Length, r.Breadth)
}

// Empty Interface : it can hold any type of value
func describe(i interface{}) {
	fmt.Printf(" Type= %T and Value= %v \n", i, i)
}

func main() {

	r1 := Rectangle{10, 20}
	c1 := Circle{10}
	fmt.Println(r1) // it will call String()
	fmt.Println(c1)
	fmt.Printf("Area of Rectangle is %v \n", r1.Area())
	fmt.Printf("Area of Circle is %f \n", c1.Area())

	var d1 float64 = 11.111
	fmt.Println(d1)
	fmt.Println(int(d1))

	fmt.Println(("*****************************"))

	var h1 interface{}
	fmt.Println(h1)

	h1 = 22
	fmt.Printf(" the type of h1= %T and value= %v \n", h1, h1)

	h1 = "mitesh"
	fmt.Printf(" the type of h1= %T and value= %v \n", h1, h1)

	x := 11
	describe(x)
	y := "mayur"
	describe(y)

	type strt struct {
		name string
	}
	z := strt{"mayur"}
	describe(z)

	fmt.Println("****************")

	var m interface{} = "himanshu"

	var s string = m.(string)
	fmt.Println(s)

	s, ok := m.(string)
	fmt.Println(s, ok)

	f, ok := m.(float64)
	fmt.Println(f, ok)

	// f = m.(float64) // 	Panic
	// fmt.Println(f)

}
