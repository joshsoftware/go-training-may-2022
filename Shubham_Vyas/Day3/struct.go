package main

import (
	"fmt"
	"time"
)

type Attendee struct {
	Name    string
	Country string
	phone   string
	ticket  string
}

type Speaker struct {
	// inheritance
	Attendee
	slot time.Time
}

type Conference struct {
	Name     string
	Location string
	people   *[]Attendee
}

type User struct {
	FirstName string
	LastName  string
	Address
}

type Address struct {
	AddressLine1 string
	Country      string
}

func learnStruct() {
	attendee := Attendee{
		Name:    "Shubham",
		Country: "India",
		phone:   "8299000000",
		ticket:  "D-1",
	}

	fmt.Println(attendee)

	speaker := Speaker{}
	// attendee properties can directly be used when no name is given
	speaker.Name = "Shubham"

	fmt.Println(speaker)

	attendees := []Attendee{attendee}
	conference := Conference{
		people: &attendees,
	}
	fmt.Println(conference)
}

func learnStruct1() {
	adr := Address{
		AddressLine1: "Line 1",
		Country:      "India",
	}
	x := User{
		FirstName: "Shubham",
		LastName:  "Vyas",
		Address:   adr,
	}

	fmt.Println(x.Address.AddressLine1)
}
