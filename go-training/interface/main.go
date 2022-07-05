package main

import (
	"fmt"
	"math"
)

// interface
type geometry interface {
	area() float64
	perim() float64
}

//strcuts : rectangle
type rectangle struct {
	height, width float64
}

// circle
type circle struct {
	radius float64
}

//methods
func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perim() float64 {
	return 2 * r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//function for interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	//values to structs
	r1 := rectangle{6, 8}
	c1 := circle{8}

	//calling functions
	measure(r1)
	measure(c1)
}
