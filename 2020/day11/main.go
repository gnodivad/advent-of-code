package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	seats := readStringsFromFile("2020/day11/input.txt")

	fmt.Println("--- Part One ---")
	fmt.Println(getOccupiedSeatWithExactAdjacentSeats(seats))

	fmt.Println("--- Part Two ---")
	fmt.Println(getOccupiedSeatWithNotExactAdjacentSeats(seats))
}

func getOccupiedSeatWithNotExactAdjacentSeats(seats [][]string) int {

	for {
		cloneSeats := cloneSeat(seats)
		isSeatChanged := false

		for rNum := range seats {
			for cNum, seat := range seats[rNum] {
				if seat == "." {
					continue
				}

				adjacentSeatCount := getNotExactAdjacentSeatCount(seats, rNum, cNum)

				if seats[rNum][cNum] == "L" && adjacentSeatCount == 0 {
					cloneSeats[rNum][cNum] = "#"
				} else if seats[rNum][cNum] == "#" && adjacentSeatCount >= 5 {
					cloneSeats[rNum][cNum] = "L"
				}

				if seats[rNum][cNum] != cloneSeats[rNum][cNum] {
					isSeatChanged = true
				}
			}
		}

		seats = cloneSeats

		if !isSeatChanged {
			break
		}
	}

	return calculateOccupiedSeat(seats)
}

func getNotExactAdjacentSeatCount(seats [][]string, rNum int, cNum int) int {
	adjacentDirection := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	count := 0
	for _, direction := range adjacentDirection {
		row, col := rNum, cNum
		for {
			row, col = row+direction[0], col+direction[1]

			if row < 0 || row >= len(seats) || col < 0 || col >= len(seats[0]) {
				break
			}

			if seats[row][col] == "." {
				continue
			}

			if seats[row][col] == "L" {
				break
			}

			if seats[row][col] == "#" {
				count++
				break
			}
		}
	}

	return count
}

func getOccupiedSeatWithExactAdjacentSeats(seats [][]string) int {
	for {
		cloneSeats := cloneSeat(seats)
		isSeatChanged := false

		for rNum := range seats {
			for cNum, seat := range seats[rNum] {
				if seat == "." {
					continue
				}

				adjacentSeatCount := getAdjacentSeatCount(seats, rNum, cNum)
				if seats[rNum][cNum] == "L" && adjacentSeatCount == 0 {
					cloneSeats[rNum][cNum] = "#"
				} else if seats[rNum][cNum] == "#" && adjacentSeatCount >= 4 {
					cloneSeats[rNum][cNum] = "L"
				}

				if seats[rNum][cNum] != cloneSeats[rNum][cNum] {
					isSeatChanged = true
				}
			}
		}

		seats = cloneSeats

		if !isSeatChanged {
			break
		}
	}

	return calculateOccupiedSeat(seats)
}

func cloneSeat(seats [][]string) [][]string {
	cloneSeats := make([][]string, len(seats))
	for i := range seats {
		cloneSeats[i] = make([]string, len(seats[i]))
		copy(cloneSeats[i], seats[i])
	}

	return cloneSeats
}

func calculateOccupiedSeat(seats [][]string) (count int) {
	count = 0
	for rNum := range seats {
		for _, seat := range seats[rNum] {
			if seat == "#" {
				count++
			}
		}
	}

	return
}

func getAdjacentSeatCount(seats [][]string, rNum int, cNum int) int {
	adjacentDirection := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	count := 0
	for _, direction := range adjacentDirection {
		row, col := rNum+direction[0], cNum+direction[1]

		if row < 0 || row >= len(seats) || col < 0 || col >= len(seats[0]) {
			continue
		}

		if seats[row][col] == "." {
			continue
		}

		if seats[row][col] == "#" {
			count++
		}
	}

	return count
}

func readStringsFromFile(filepath string) (seats [][]string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		line := make([]string, 0, len(s))

		for _, c := range s {
			line = append(line, string(c))
		}

		seats = append(seats, line)
	}

	return
}
