package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	numbers := readNumbersFromFile("2020/day10/input.txt")
	numbers = append(numbers, 0)
	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	fmt.Println("--- Part One ---")
	fmt.Println(getTheMultipyOf1JoltDiffWith3JoltDiff(numbers))

	memo := make([]int, len(numbers), len(numbers))
	for i := range memo {
		memo[i] = -1
	}
	fmt.Println("--- Part Two ---")
	fmt.Println(getTotalOfArrangement(numbers, memo, 0))
}

func getTheMultipyOf1JoltDiffWith3JoltDiff(numbers []int) int {
	diffWith1, diffWith3 := 0, 0
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		if diff == 1 {
			diffWith1++
		} else if diff == 3 {
			diffWith3++
		}
	}

	return diffWith1 * diffWith3
}

func getTotalOfArrangement(numbers []int, memo []int, index int) int {
	if index == len(numbers)-1 {
		return 1
	}

	if index >= len(numbers) {
		return 0
	}

	if memo[index] == -1 {
		count := 0
		for i, number := range numbers[index+1:] {
			if number-numbers[index] > 3 {
				break
			}
			count += getTotalOfArrangement(numbers, memo, (index + 1 + i))
		}
		memo[index] = count
	}

	return memo[index]
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
