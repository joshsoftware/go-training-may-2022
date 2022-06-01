package interfacesexample

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
	Name   string
}
type Circle struct {
	Redius float64
}

type areaCalculate interface {
	area() float64
}

func (s Square) area() float64 {
	return s.Length * s.Length
}

func (c Circle) area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func PrintAreas(a areaCalculate) {
	fmt.Println(a.area())
}

func main() {
	s := Square{20, "fas"}
	c := Circle{20}
	PrintAreas(s)
	PrintAreas(c)
}
