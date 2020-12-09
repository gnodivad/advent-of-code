package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bagCount struct {
	name  string
	count int
}

const shinyGold = "shiny gold"

func main() {
	ruleRows := readStringsFromFile("day07/input.txt")

	contains := make(map[string][]bagCount)

	for _, ruleRow := range ruleRows {
		ruleRow = strings.TrimSuffix(ruleRow, ".")
		ruleRowSlice := strings.Split(ruleRow, " contain ")
		bagType, containsList := ruleRowSlice[0], strings.Split(ruleRowSlice[1], ", ")

		bagType = removeBagsWordFromString(bagType)

		var bags []bagCount
		for _, containsBag := range containsList {
			containsBagSlice := strings.SplitAfterN(containsBag, " ", 2)
			if containsBagSlice[0] != "no" {
				c, _ := strconv.Atoi(strings.TrimSpace(containsBagSlice[0]))
				bags = append(bags, bagCount{removeBagsWordFromString(containsBagSlice[1]), c})
			}
		}
		contains[bagType] = bags
	}

	count := 0

	for _, v := range contains {
		for _, containItem := range v {
			if findColor(containItem.name, contains) {
				count++
				break
			}
		}
	}

	var countBags func(b string) int
	countBags = func(b string) int {
		var count int
		for _, cb := range contains[b] {
			count += cb.count * (countBags(cb.name) + 1)
		}
		return count
	}

	fmt.Println("--- Part One ---")
	fmt.Println(count)

	fmt.Println("--- Part Two ---")
	fmt.Println(countBags(shinyGold))
}

func findColor(s string, contains map[string][]bagCount) bool {
	if s == shinyGold {
		return true
	}

	isColorFind := false

	if val, ok := contains[s]; ok {
		for _, containItem := range val {
			if findColor(containItem.name, contains) {
				isColorFind = true
				break
			}
		}
	}
	return isColorFind
}

func removeBagsWordFromString(s string) string {
	return strings.TrimSuffix(strings.TrimSuffix(s, "s"), " bag")
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
