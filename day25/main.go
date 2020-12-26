package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
)

func main() {
	input := utils.ReadStringsFromFile("day25/input.txt")
	cardPublicKey := utils.ParseInt(input[0])
	doorPublicKey := utils.ParseInt(input[1])

	fmt.Println("--- Part One ---")
	fmt.Println(transformSubjectNumber(doorPublicKey, findLoopSize(cardPublicKey)))
}

func transformSubjectNumber(subjectNumber int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = (value * subjectNumber) % 20201227
	}

	return value
}

func findLoopSize(cardPublicKey int) int {
	i, value := 1, 1

	for {
		value = (value * 7) % 20201227
		if value == cardPublicKey {
			return i
		}

		i++
	}
}
