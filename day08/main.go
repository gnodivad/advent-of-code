package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := readStringsFromFile("day08/input.txt")

	acc := 0
	fmt.Println("--- Part One ---")
	acc, _ = runMachine(instructions)
	fmt.Println(acc)

	fmt.Println("--- Part Two ---")
	acc, _ = findBug(instructions)
	fmt.Println(acc)
}

func runMachine(instructions []string) (int, bool) {
	executionPointer, acc := 0, 0
	memory := make([]bool, len(instructions), len(instructions))

	for true {
		if executionPointer == len(instructions) {
			break
		}

		instructionSlice := strings.Split(instructions[executionPointer], " ")
		operation, argumentString := instructionSlice[0], instructionSlice[1]

		if memory[executionPointer] == true {
			return acc, false
		}

		memory[executionPointer] = true

		argument, _ := strconv.Atoi(argumentString)

		if operation == "nop" {
			executionPointer++
		} else if operation == "acc" {
			acc += argument
			executionPointer++
		} else if operation == "jmp" {
			executionPointer += argument
		}
	}

	return acc, true
}

func findBug(instructions []string) (int, bool) {
	for i, instruction := range instructions {
		instructionSlice := strings.Split(instruction, " ")
		operation := instructionSlice[0]

		var modifiedInstruction []string
		if operation == "nop" {
			modifiedInstruction = cloneInstruction(instructions)
			modifiedInstruction[i] = strings.Replace(modifiedInstruction[i], "nop", "jmp", 1)
		} else if operation == "jmp" {
			modifiedInstruction = cloneInstruction(instructions)
			modifiedInstruction[i] = strings.Replace(modifiedInstruction[i], "jmp", "nop", 1)
		}

		if modifiedInstruction != nil {
			acc, done := runMachine(modifiedInstruction)

			if done {
				return acc, done
			}
		}
	}

	return 0, false
}

func cloneInstruction(instruction []string) []string {
	modifiedInstruction := make([]string, len(instruction))
	copy(modifiedInstruction, instruction)
	return modifiedInstruction
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
