package main

import "fmt"

type person struct {
	name       string
	speaking   string
	compayInfo company
}
type company struct {
	name     string
	location string
}

func (p person) pesonalInfo() {
	p.compayInfo.name = "Josh"
	fmt.Println(p.name, "Speaks", p.speaking, "works in", p.compayInfo.name)
}

func personalInfo(p *person) {
	p.compayInfo.name = "xyz"
	fmt.Println(p.name, "Speaks", p.speaking, "works in", p.compayInfo.name)
}

func Receiverfunctionexample() {
	p1 := person{
		name:     "Akshay",
		speaking: "Marathi",
	}
	p2 := person{
		name:     "Leo",
		speaking: "English",
		compayInfo: company{
			name:     "Sapiens",
			location: "UK",
		},
	}
	fmt.Println(p1, p2)
	fmt.Println("-----------Output in personalInfo method--------------------")
	p1.pesonalInfo()
	p2.pesonalInfo()
	fmt.Println("-----------Output in Receiverfunctionexample method---------")
	fmt.Println(p1, p2)
	//personalInfo(&p1)
	//personalInfo(&p2)
	fmt.Println("----------------------------")

	foo()
}

func foo() {
	a := 4
	b := 6
	defer func() {
		fmt.Println(a + b)
	}()

	//b = 5

	//defer fmt.Println(a + b)
	return
}

func main() {
	Receiverfunctionexample()
}
