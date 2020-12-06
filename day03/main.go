package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	maps := readStringsFromFile("day03/input.txt")

	fmt.Println("--- Part One ---")
	fmt.Println(traverseMap(maps, 1, 3))

	fmt.Println("--- Part Two ---")
	fmt.Println(traverseMap(maps, 1, 1) * traverseMap(maps, 1, 3) * traverseMap(maps, 1, 5) * traverseMap(maps, 1, 7) * traverseMap(maps, 2, 1))
}

func traverseMap(maps [][]bool, moveDown int, moveRight int) (treeEncoutered int) {
	row := 0 + moveDown
	col := 0 + moveRight

	for row < len(maps) {
		if maps[row][col] {
			treeEncoutered++
		}
		col = (col + moveRight) % len(maps[0])
		row += moveDown
	}

	return
}

func readStringsFromFile(filepath string) (booleans [][]bool) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		line := make([]bool, 0, len(s))

		for _, c := range s {
			line = append(line, string(c) == "#")
		}

		booleans = append(booleans, line)
	}

	return
}
