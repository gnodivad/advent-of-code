package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const threshold int = 25

func main() {
	numbers := readNumbersFromFile("day09/input.txt")

	fmt.Println("--- Part One ---")
	invalidTwoSum := slidingTwoSum(numbers)
	fmt.Println(invalidTwoSum)
	fmt.Println(slidingSum(numbers, invalidTwoSum))
}

func slidingSum(numbers []int, target int) int {
	left, right := 0, 1
	sum := numbers[left] + numbers[right]

	for right < len(numbers) {
		if sum == target {
			break
		} else if sum < target {
			right++
			sum += numbers[right]
		} else {
			sum -= numbers[left]
			left++
		}
	}

	return findTotalOfMixMax(numbers[left : right+1])
}

func findTotalOfMixMax(numbers []int) int {
	var min, max int

	for _, number := range numbers {
		if number > max {
			max = number
		}

		if min == 0 || number < min {
			min = number
		}
	}

	return min + max
}

func slidingTwoSum(numbers []int) int {
	for i, number := range numbers {
		if i >= threshold {
			if !isTwoSum(numbers[i-threshold:i], number) {
				return number
			}
		}
	}

	return -1
}

func isTwoSum(numbers []int, target int) bool {
	lookup := make(map[int]bool)

	for _, v := range numbers {
		if lookup[v] {
			return true
		}
		lookup[target-v] = true
	}

	return false
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
