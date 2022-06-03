package main

import (
	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/bwmarrin/discordgo"
)

func main() {
	DescordModule()
	RandomData()
}

// DescordModule function will generate the discord session object
func DescordModule() {
	discord, err := discordgo.New("")
	if err != nil {
		fmt.Println("Could not start the session")
	}
	fmt.Println("Session created: ", discord)
}

//RandomData generates the random data
func RandomData() {
	// Print a random silly name
	fmt.Println(randomdata.SillyName())

	// Print a male title
	fmt.Println(randomdata.Title(randomdata.Male))

	// Print a female title
	fmt.Println(randomdata.Title(randomdata.Female))
}
