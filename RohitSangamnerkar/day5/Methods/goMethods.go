package main

import "fmt"

//methods in golang are similer to golang function
//the diffrence is it has reciver argument
//using that reciver method can acess recivers fields

/*
	func(reciver_name Type) method_name(parameter_list)(return_type){
		// Code
	}
*/

//create one structure
type person struct {
	id        int
	firstName string
	lastName  string
	age       int
}

//creating a methods with reciver type person
//this is value type reciver so we need to changed person obj
//if we change value of person it will not reflect in original struct

func (p person) getPerson(id int) person {
	if p.id == id {
		p.firstName = "Rohit"
		p.lastName = "Sangamnerkar"
		p.age = 23

		return p
	} else {
		return p
	}
}

//if we want to change the value of fields of reciver then
//we need to use pointer reciver

func (p *person) setInfoOfPerson(id int, fname, lname string, age int) string {
	if p.id == id {
		p.firstName = fname
		p.lastName = lname
		p.age = age

		return "employee details updated sucessfully"
	}

	return "employee not found"
}

func main() {

	var personObj person

	personObj.id = 10
	//personObj.getPerson(11)
	recivedPerson := personObj.getPerson(10)

	fmt.Println(recivedPerson)

	//create another obj of struct

	var personObj2 person
	personObj2.id = 22

	//print person =Obj2 before updating its info
	fmt.Println("Before updation : ", personObj2)
	//call setInfo function to setInfo
	msg := personObj2.setInfoOfPerson(22, "Rohit", "Sharma", 27)
	fmt.Println(msg)
	fmt.Println("after updation : ", personObj2)
}
