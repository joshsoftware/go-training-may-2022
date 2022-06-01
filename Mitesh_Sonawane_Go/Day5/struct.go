package main

import "fmt"

type Student struct {
	rollno int
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func main() {

	fmt.Println(Student{})
	s1 := Student{11, "mitesh", []int{10, 20, 30}, 24}
	fmt.Println(s1)
	fmt.Printf("student rollno is %v and name is %v \n", s1.rollno, s1.name)
	fmt.Println("student age is ", s1.getAge())
	s1.setAge(33)
	fmt.Println(s1)
	fmt.Println(Student{55, "mayur", []int{40, 50}, 19})

}
