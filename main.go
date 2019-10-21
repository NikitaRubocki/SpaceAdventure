package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "io/ioutil"
	// "encoding/json"
)

// func makePlanets()

func main() {
	// variable declarations
	var name = "Solar System"
	planets := map[string]string{
		"Mercury": "A very hot planet, closest to the sun.",
		"Venus": "It's very cloudy here!",
		"Earth": "There is something very familiar about this planet.",
		"Mars": "Known as the red planet.",
		"Jupiter": "A gas giant, with a noticeable red spot.", 
		"Saturn": "This planet has beautiful rings around it.",
		"Uranus": "Strangely, this planet rotates around on its side.",
		"Neptune": "A very cold planet, furthest from the sun.",
		"Pluto": "I don't care what they say - it's a planet.",
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Welcome to the %v!\n", name)
	fmt.Printf("There are %v planets to explore.\n", len(planets))
	fmt.Println("What is your name?")
	//var userName string
	input, _ := reader.ReadString('\n')
	userName := strings.TrimSuffix(input, "\n")
	fmt.Println("Nice to meet you, "+userName+". My name is Eliza, I'm an old friend of Alexa.")
	
	// fmt.Println(planets["Earth"])
}

// file, _ := ioutil.ReadFile("planetarySystem.json")
// planets := json.Unmarshal(file, &Space)
// m := make(map[string]string)
	// type System struct {
	// 	Name string
	// 	Planets []Planet
	// }
	//var Space map[string]Planet
	// var Space System
	// type Planet []struct{
	// 	Name, Description string
	// }