package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Planet name and description
type Planet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// System name and arrays of Planets
type System struct {
	Name    string   `json:"name"`
	Planets []Planet `json:"planets"`
}

func makeSolarSystem(file string) System {
	jsonFile, _ := os.Open(file)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var solarSystem System
	json.Unmarshal(byteValue, &solarSystem)
	return solarSystem
}

func validateInput(answer string, reader *bufio.Reader) string {
	if answer == "N" || answer == "Y" {
		return answer
	}
	fmt.Println("Sorry, I didn't get that.")
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
	answer, _ = reader.ReadString('\n')
	return validateInput(strings.ToUpper(strings.TrimSuffix(answer, "\n")), reader)
}

func validatePlanet(answer string, planets []Planet, reader *bufio.Reader) Planet {
	for i := 0; i < len(planets); i++ {
		if answer == planets[i].Name {
			return planets[i]
		}
	}
	fmt.Println("Sorry, that planet doesn't exist.")
	fmt.Println("Name the planet you would like to visit.")
	choice, _ := reader.ReadString('\n')
	return validatePlanet(strings.TrimSuffix(choice, "\n"), planets, reader)
}

func choosePlanet(input string, solarSystem System, reader *bufio.Reader, r *rand.Rand) {
	if input == "N" {
		fmt.Println("Name the planet you would like to visit.")
		choice, _ := reader.ReadString('\n')
		planet := validatePlanet(strings.TrimSuffix(choice, "\n"), solarSystem.Planets, reader)
		fmt.Printf("Traveling to %v...\n", planet.Name)
		fmt.Printf("Arrived at %v! %v\n", planet.Name, planet.Description)
	} else {
		num := r.Intn(len(solarSystem.Planets))
		fmt.Printf("Traveling to %v...\n", solarSystem.Planets[num].Name)
		fmt.Printf("Arrived at %v! %v\n", solarSystem.Planets[num].Name, solarSystem.Planets[num].Description)
	}
}

func main() {
	
	// variable declarations
	reader := bufio.NewReader(os.Stdin)
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// introduction
	solarSystem := makeSolarSystem("planetarySystem.json")
	fmt.Printf("Welcome to the %v!\n", solarSystem.Name)
	fmt.Printf("There are %v planets to explore.\n", len(solarSystem.Planets))

	// get username and print
	fmt.Println("What is your name?")
	input, _ := reader.ReadString('\n')
	userName := strings.TrimSuffix(input, "\n")
	fmt.Printf("Nice to meet you, %v. My name is Eliza, I'm an old friend of Alexa.", userName)

	// the adventure begins
	fmt.Println("Let's go on an adventure!")
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
	answer, _ := reader.ReadString('\n')

	// validate user input and choose a planet
	valAnswer := validateInput(strings.ToUpper(strings.TrimSuffix(answer, "\n")), reader)
	choosePlanet(valAnswer, solarSystem, reader, r1)

}