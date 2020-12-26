package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var myExp = regexp.MustCompile(`(?P<action>\D+)(?P<units>\d+)`)

type Instruction struct {
	action string
	units  int
}

type Coordinate struct {
	x int
	y int
}

type Spaceship struct {
	towardDirection Coordinate
	coordinate      Coordinate
}

func (c Coordinate) turn(d int) Coordinate {
	dp := float64(d) * math.Pi / 180
	return Coordinate{
		x: int(math.Round(math.Cos(dp)*float64(c.x) - math.Sin(dp)*float64(c.y))),
		y: int(math.Round(math.Sin(dp)*float64(c.x) + math.Cos(dp)*float64(c.y))),
	}
}

func (spaceship *Spaceship) execute(instruction Instruction) {
	if instruction.action == "N" {
		spaceship.coordinate.y += instruction.units
	} else if instruction.action == "S" {
		spaceship.coordinate.y -= instruction.units
	} else if instruction.action == "E" {
		spaceship.coordinate.x += instruction.units
	} else if instruction.action == "W" {
		spaceship.coordinate.x -= instruction.units
	} else if instruction.action == "L" {
		spaceship.towardDirection = spaceship.towardDirection.turn(instruction.units)
	} else if instruction.action == "R" {
		spaceship.towardDirection = spaceship.towardDirection.turn(-instruction.units)
	} else if instruction.action == "F" {
		spaceship.coordinate.x = spaceship.coordinate.x + (instruction.units * spaceship.towardDirection.x)
		spaceship.coordinate.y = spaceship.coordinate.y + (instruction.units * spaceship.towardDirection.y)
	}
}

func (spaceship *Spaceship) executeAlongWithWaypoint(instruction Instruction) {
	if instruction.action == "N" {
		spaceship.towardDirection.y += instruction.units
	} else if instruction.action == "S" {
		spaceship.towardDirection.y -= instruction.units
	} else if instruction.action == "E" {
		spaceship.towardDirection.x += instruction.units
	} else if instruction.action == "W" {
		spaceship.towardDirection.x -= instruction.units
	} else if instruction.action == "L" {
		spaceship.towardDirection = spaceship.towardDirection.turn(instruction.units)
	} else if instruction.action == "R" {
		spaceship.towardDirection = spaceship.towardDirection.turn(-instruction.units)
	} else if instruction.action == "F" {
		spaceship.coordinate.x = spaceship.coordinate.x + (instruction.units * spaceship.towardDirection.x)
		spaceship.coordinate.y = spaceship.coordinate.y + (instruction.units * spaceship.towardDirection.y)
	}
}

func main() {
	instructionTexts := readStringsFromFile("2020/day12/input.txt")
	instructions := parseInstructions(instructionTexts)

	spaceship1 := Spaceship{towardDirection: Coordinate{x: 1, y: 0}, coordinate: Coordinate{x: 0, y: 0}}
	spaceship2 := Spaceship{towardDirection: Coordinate{x: 10, y: 1}, coordinate: Coordinate{x: 0, y: 0}}

	for _, instruction := range instructions {
		spaceship1.execute(instruction)
		spaceship2.executeAlongWithWaypoint(instruction)
	}

	fmt.Println("--- Part One ---")
	fmt.Println(abs(spaceship1.coordinate.x) + abs(spaceship1.coordinate.y))

	fmt.Println("--- Part Two ---")
	fmt.Println(abs(spaceship2.coordinate.x) + abs(spaceship2.coordinate.y))
}

func parseInstructions(instructionTexts []string) []Instruction {
	instructions := make([]Instruction, len(instructionTexts))
	for index, instructionText := range instructionTexts {
		match := myExp.FindStringSubmatch(instructionText)
		result := make(map[string]string)
		for i, name := range myExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		units, _ := strconv.ParseInt(result["units"], 10, 0)
		instructions[index] = Instruction{
			action: result["action"],
			units:  int(units),
		}
	}

	return instructions
}

func abs(x int) int {
	if x < 0 {
		return -(x)
	}

	return x
}

func readStringsFromFile(filepath string) (strings []string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return
}
