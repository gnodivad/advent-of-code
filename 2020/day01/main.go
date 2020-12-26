package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	expenses := readNumbersFromFile("2020/day01/input.txt")

	fmt.Println("--- Part One ---")
	fmt.Println(productOfTwoSumNumber(expenses))

	fmt.Println("--- Part Two ---")
	fmt.Println(productOfThreeSumNumber(expenses))
}

func productOfTwoSumNumber(expenses []int) int {
	lookupTable := make(map[int]bool, len(expenses))

	for _, expense := range expenses {
		if lookupTable[2020-expense] {
			return expense * (2020 - expense)
		}

		lookupTable[expense] = true
	}

	return 0
}

func productOfThreeSumNumber(expenses []int) int {
	lookupTable := make(map[int]bool, len(expenses))

	for _, expense := range expenses {
		lookupTable[expense] = true
	}

	for i, expense := range expenses {
		for _, secondExpense := range expenses[i+1:] {
			if lookupTable[2020-expense-secondExpense] {
				return expense * secondExpense * (2020 - expense - secondExpense)
			}
		}
	}

	return 0
}

func readNumbersFromFile(filepath string) (numbers []int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = append(numbers, stringToInt(scanner.Text()))
	}

	return
}

func stringToInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
