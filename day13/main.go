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
	notes := readStringsFromFile("day13/input.txt")
	departAtStr, busList := notes[0], notes[1]

	departAt, _ := strconv.Atoi(departAtStr)

	var buses []int
	for _, busStr := range strings.Split(busList, ",") {
		if busStr == "x" {
			buses = append(buses, -1)
		} else {
			busAsInt, _ := strconv.Atoi(busStr)
			buses = append(buses, busAsInt)
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(findBus(departAt, buses))

	fmt.Println("--- Part Two ---")
	fmt.Println(findStartOfBusSubsequentDepartedTimeSeries(buses))
}

func findBus(departAt int, buses []int) int {
	busSchedule := make(map[int]int)

	for _, bus := range buses {
		if bus == -1 {
			continue
		}

		nextDepartAt := departAt
		for {
			if nextDepartAt%bus == 0 {
				busSchedule[bus] = nextDepartAt
				break
			}
			nextDepartAt++
		}
	}

	minBus := 0
	for k, v := range busSchedule {
		if minBus == 0 {
			minBus = k
		}

		if busSchedule[minBus]-departAt > v-departAt {
			minBus = k
		}
	}

	return minBus * (busSchedule[minBus] - departAt)
}

// Reference to https://www.mathsisfun.com/least-common-multiple.html
func findStartOfBusSubsequentDepartedTimeSeries(buses []int) int {
	start := 0
	isSeries := false

	for !isSeries {
		lcm := 1
		isSeries, lcm = findLeastCommonMultiple(buses, start, lcm)

		if !isSeries {
			start = start + lcm
		}
	}

	return start
}

func findLeastCommonMultiple(buses []int, start int, skipFactor int) (bool, int) {
	for i, bus := range buses {
		if bus == -1 {
			continue
		}

		if (start+i)%bus != 0 {
			return false, skipFactor
		}

		skipFactor *= bus
	}

	return true, skipFactor
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
