package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	answerRows := readStringsFromFile("day06/input.txt")

	questions := make(map[string]int)
	count := 0
	peopleCount := 0
	everyoneCount := 0
	for _, answerRow := range answerRows {
		if answerRow == "" {
			for _, v := range questions {
				count++
				if v == peopleCount {
					everyoneCount++
				}
			}
			questions = make(map[string]int)
			peopleCount = 0
			continue
		}

		peopleCount++

		for _, answer := range answerRow {
			questions[string(answer)] = questions[string(answer)] + 1
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(count)

	fmt.Println("--- Part Two ---")
	fmt.Println(everyoneCount)
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
