package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	boardingPasses := readStringsFromFile("day05/input.txt")

	minSeatId, maxSeatId := 128*8, -1

	var seats [128 * 8]bool

	for _, boardingPass := range boardingPasses {
		seatId := getSeatId(boardingPass)

		minSeatId = min(minSeatId, seatId)
		maxSeatId = max(maxSeatId, seatId)

		seats[seatId] = true

	}

	emptySeat := 0
	for id, used := range seats[minSeatId+1 : maxSeatId] {
		if !used {
			emptySeat = minSeatId + 1 + id
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(maxSeatId)

	fmt.Println("--- Part Two ---")
	fmt.Println(emptySeat)
}

func getSeatId(boardingPass string) int {
	rowList := getList(128)

	left, mid, right := rowList[0], rowList[0], rowList[len(rowList)-1]
	for _, char := range boardingPass[0:7] {
		mid = left + (right-left)/2

		if char == 'F' {
			right = mid
		} else if char == 'B' {
			left = mid + 1
		}
	}
	row := rowList[left]

	columnList := getList(8)
	left, mid, right = columnList[0], columnList[0], columnList[len(columnList)-1]
	for _, char := range boardingPass[7:] {
		mid = left + (right-left)/2

		if char == 'L' {
			right = mid
		} else if char == 'R' {
			left = mid + 1
		}
	}
	column := columnList[left]

	return row*8 + column
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}

func getList(n int) []int {
	list := make([]int, n)
	for i := range list {
		list[i] = i
	}
	return list
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
