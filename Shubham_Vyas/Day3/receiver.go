package main

import "fmt"

type Car struct {
	Color string
	Owner string
	ID    int
}

// Now instead of passing car obj using arguments we add a method to Car struct
// When using * (pointers) we can pass data using reference
// If no * is used it is passed as value
func (c *Car) updateCarDetails() {
	if c.ID == 1 {
		c.Owner = "Shubham"
		c.Color = "Black"
	} else {
		c.Owner = "Josh"
		c.Color = "White"
	}
}

func carDetails(c Car) {
	c.Color = "White"
	c.Owner = "John"

	fmt.Printf("Updated using value type - %s owns a %s car\n", c.Owner, c.Color)
}

func carDetails2(c *Car) {
	c.Color = "White"
	c.Owner = "John"

	fmt.Printf("Updated using ref type - %s owns a %s car\n", c.Owner, c.Color)
}
func learnReceivers() {
	c := Car{
		ID: 1,
	}

	carDetails(c)
	// updated fields wont be reflected as we passed it by value
	fmt.Printf("Updated car details: %s owns a %s car\n", c.Owner, c.Color)

	// Lets pass it by reference in receiver
	carDetails2(&c)
	fmt.Printf("Updated car details: %s owns a %s car\n", c.Owner, c.Color)

	// As you can see the changes are reflected

	// Its same using methods
	c.updateCarDetails()
	newCar := Car{
		ID: 2,
	}
	newCar.updateCarDetails()
	fmt.Printf("Updated car details: %s owns a %s car\n", c.Owner, c.Color)
	fmt.Printf("Updated car details: %s owns a %s car\n", newCar.Owner, newCar.Color)
}
