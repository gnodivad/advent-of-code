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
	numbersAsString := readStringsFromFile("day15/input.txt")
	numbers := make([]int64, 30000000)

	len := int64(0)
	lookup := make(map[int64]int64)
	for _, numberAsString := range strings.Split(numbersAsString[0], ",") {
		numberAsInt, _ := strconv.Atoi(numberAsString)
		numbers[len] = int64(numberAsInt)
		lookup[int64(numberAsInt)] = len
		len++
	}

	for i := range numbers[len:] {
		currentIndex := len + int64(i)
		prevNumber := numbers[currentIndex-1]

		if _, ok := lookup[prevNumber]; !ok {
			numbers[currentIndex] = 0
			lookup[prevNumber] = currentIndex - 1
		} else {
			numbers[currentIndex] = currentIndex - 1 - lookup[prevNumber]
			lookup[prevNumber] = currentIndex - 1
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(numbers[2019])

	fmt.Println("--- Part Two ---")
	fmt.Println(numbers[29999999])
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
