package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"time"
	"unicode"
	"io/ioutil"
	"encoding/json"
)

// Planet name and description
type Planet struct {
	Name string 		`json:"name"`
	Description string	`json:"description"`
}

// System name and arrays of Planets
type System struct {
	Name string 		`json:"name"`
	Planets []Planet	`json:"planets"`
}

func makeSolarSystem(file string) System {
	jsonFile, _ := os.Open(file)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var solarSystem System
	json.Unmarshal(byteValue, &solarSystem)
	return solarSystem
}

func validateYorN(answer rune) rune {
	runeReader := bufio.NewReader(os.Stdin)
	if answer == 'N' || answer == 'Y'{
		return answer
	}
	fmt.Println("Sorry, I didn't get that.")
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
	answer, _, _ = runeReader.ReadRune()
	return validateYorN(unicode.ToUpper(answer))
}

func validatePlanet(answer string, planets []Planet) Planet {
	stringReader := bufio.NewReader(os.Stdin)
	for i:=0; i<len(planets); i++{
		if answer == planets[i].Name{
			return planets[i]
		}
	}
	fmt.Println("Sorry, that planet doesn't exist.")
	fmt.Println("Name the planet you would like to visit.")
	choice, _ := stringReader.ReadString('\n')
	return validatePlanet(strings.TrimSuffix(choice, "\n"), planets)
}

func choosePlanet(input rune, solarSystem System, r *rand.Rand){
	stringReader := bufio.NewReader(os.Stdin)
	if input == 'N'{
		fmt.Println("Name the planet you would like to visit.")
		choice, _ := stringReader.ReadString('\n')
		planet := validatePlanet(strings.TrimSuffix(choice, "\n"), solarSystem.Planets)
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
	stringReader := bufio.NewReader(os.Stdin)
	runeReader := bufio.NewReader(os.Stdin)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// var name = "Solar System"
	// planets := map[string]string{
	// 	"Mercury": "A very hot planet, closest to the sun.",
	// 	"Venus": "It's very cloudy here!",
	// 	"Earth": "There is something very familiar about this planet.",
	// 	"Mars": "Known as the red planet.",
	// 	"Jupiter": "A gas giant, with a noticeable red spot.", 
	// 	"Saturn": "This planet has beautiful rings around it.",
	// 	"Uranus": "Strangely, this planet rotates around on its side.",
	// 	"Neptune": "A very cold planet, furthest from the sun.",
	// 	"Pluto": "I don't care what they say - it's a planet.",
	// }
    // keys := make([]string, 0, len(planets))
    // for k := range planets {
    //     keys = append(keys, k)
	// }

	// type Planet struct{
	// 	Name string 		`json:"name"`
	// 	Description string	`json:"description"`
	// }
	// type System struct {
	// 	Name string 		`json:"name"`
	// 	Planets []Planet	`json:"planets"`
	// }

	// working with json file
	// jsonFile, _ := os.Open("planetarySystem.json")
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// var solarSystem System
	// json.Unmarshal(byteValue, &solarSystem)

	// introduction
	solarSystem := makeSolarSystem("planetarySystem.json")
	fmt.Printf("Welcome to the %v!\n", solarSystem.Name)
	fmt.Printf("There are %v planets to explore.\n", len(solarSystem.Planets))

	// get username and print
	fmt.Println("What is your name?")
	input, _ := stringReader.ReadString('\n')
	userName := strings.TrimSuffix(input, "\n")
	fmt.Println("Nice to meet you, "+userName+". My name is Eliza, I'm an old friend of Alexa.")

	// the adventure begins
	fmt.Println("Let's go on an adventure!")
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
	answer, _, _ := runeReader.ReadRune()

	//validate user input
	upAnswer := unicode.ToUpper(answer)
	valAnswer := validateYorN(upAnswer)
	choosePlanet(valAnswer, solarSystem, r1)

}

// file, _ := ioutil.ReadFile("planetarySystem.json")
// planets := json.Unmarshal(file, &Space)
// m := make(map[string]string)

	//var Space map[string]Planet
	// var Space System


	// planets := []Planet{
	// 	{"Mercury", "A very hot planet, closest to the sun."},
	// 	{"Venus", "It's very cloudy here!"},
	// 	{"Earth", "There is something very familiar about this planet."},
	// 	{"Mars", "Known as the red planet."},
	// 	{"Jupiter", "A gas giant, with a noticeable red spot."}, 
	// 	{"Saturn", "This planet has beautiful rings around it."},
	// 	{"Uranus", "Strangely, this planet rotates around on its side."},
	// 	{"Neptune", "A very cold planet, furthest from the sun."},
	// 	{"Pluto", "I don't care what they say - it's a planet."},
	// }
