package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"time"
	"unicode"
	// "io/ioutil"
	// "encoding/json"
)

// func makePlanets()

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

func validatePlanet(answer string, keys []string) string {
	stringReader := bufio.NewReader(os.Stdin)
	for i:=0; i<len(keys); i++{
		if answer == keys[i]{
			return answer
		}
	}
	fmt.Println("Sorry, that planet doesn't exist.")
	fmt.Println("Name the planet you would like to visit.")
	choice, _ := stringReader.ReadString('\n')
	return validatePlanet(strings.TrimSuffix(choice, "\n"), keys)
}

func main() {
	// variable declarations
	stringReader := bufio.NewReader(os.Stdin)
	runeReader := bufio.NewReader(os.Stdin)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

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
    keys := make([]string, 0, len(planets))
    for k := range planets {
        keys = append(keys, k)
	}

	// introduction
	fmt.Printf("Welcome to the %v!\n", name)
	fmt.Printf("There are %v planets to explore.\n", len(planets))

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

	if valAnswer == 'N'{
		fmt.Println("Name the planet you would like to visit.")
		choice, _ := stringReader.ReadString('\n')
		planet := validatePlanet(strings.TrimSuffix(choice, "\n"), keys)
		fmt.Printf("Traveling to %v...\n", planet)
		fmt.Printf("Arrived at %v! %v\n", planet, planets[planet])
	} else {
		num := r1.Intn(len(planets))
		planet := keys[num]
		fmt.Printf("Traveling to %v...\n", planet)
		fmt.Printf("Arrived at %v! %v\n", planet, planets[planet])
	}

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

	// type Planet struct{
	// 	Name, Description string
	// }
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
