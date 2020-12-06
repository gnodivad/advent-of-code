package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := readStringsFromFile("day04/input.txt")

	list := getNewChecklist()
	validationList := getNewChecklist()
	validPasspostKeyCount := 0
	validPasspostKeyValueCount := 0
	for _, row := range passports {
		if row == "" {
			if isValidPassport(list) {
				validPasspostKeyCount++
			}

			if isValidPassport(validationList) {
				validPasspostKeyValueCount++
			}

			list = getNewChecklist()
			validationList = getNewChecklist()
			continue
		}

		for _, kvPair := range strings.Split(row, " ") {
			kvPairSlice := strings.Split(kvPair, ":")
			k, v := kvPairSlice[0], kvPairSlice[1]
			list[k] = true
			validationList[k] = isFieldContainsValidValue(k, v)
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(validPasspostKeyCount)

	fmt.Println("--- Part Two ---")
	fmt.Println(validPasspostKeyValueCount)
}

func isFieldContainsValidValue(k string, v string) bool {
	if k == "byr" {
		year, _ := strconv.Atoi(v)
		return year >= 1920 && year <= 2002
	}

	if k == "iyr" {
		year, _ := strconv.Atoi(v)
		return year >= 2010 && year <= 2020
	}

	if k == "eyr" {
		year, _ := strconv.Atoi(v)
		return year >= 2020 && year <= 2030
	}

	if k == "hgt" {
		if strings.Contains(v, "cm") {
			height, _ := strconv.Atoi(strings.ReplaceAll(v, "cm", ""))
			return height >= 150 && height <= 193
		}

		if strings.Contains(v, "in") {
			height, _ := strconv.Atoi(strings.ReplaceAll(v, "in", ""))
			return height >= 59 && height <= 76
		}
	}

	if k == "hcl" {
		var isColor = regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString
		return isColor(v)
	}

	if k == "ecl" {
		return v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "grn" || v == "hzl" || v == "oth"
	}

	if k == "pid" {
		var isNumber = regexp.MustCompile(`^\d{9}$`).MatchString
		return isNumber(v)
	}

	return false
}

func isValidPassport(list map[string]bool) (isValid bool) {
	isValid = true

	for k := range list {
		if list[k] == false && k != "cid" {
			isValid = false
		}
	}

	return
}

func getNewChecklist() (list map[string]bool) {
	list = map[string]bool{}
	list["byr"] = false
	list["iyr"] = false
	list["eyr"] = false
	list["hgt"] = false
	list["hcl"] = false
	list["ecl"] = false
	list["pid"] = false
	list["cid"] = false

	return
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
