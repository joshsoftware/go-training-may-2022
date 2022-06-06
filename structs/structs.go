package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) circumference() float64 {
	return s.length * 4
}

func (c circle) circumference() float64 {
	return math.Pi * c.radius * 2
}

func main() {
	squareObj := square{65}
	circleObj := circle{3.14}
	fmt.Println("Area of square with side ", squareObj.length, "is", squareObj.area(), "and circumference is ", squareObj.circumference())
	fmt.Println("Area of circle with radius ", circleObj.radius, "is", circleObj.area(), "and circumference is ", circleObj.circumference())

}
