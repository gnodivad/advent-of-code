package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
	"strings"
)

func main() {
	input := utils.ReadFile("2020/day20/input.txt")
	tileSet := createTileSet(input)

	result := 1
	for _, tile := range tileSet {
		if len(tile.neighbors) == 2 {
			result *= tile.id
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(result)

	fmt.Println("--- Part Two ---")
}

func createTileSet(input string) (tileSet TileSet) {
	tileSet = make(TileSet)
	tiles := strings.Split(input, "\n\n")

	for _, tileRawData := range tiles {
		s := strings.SplitN(tileRawData, "\n", 2)
		var id int
		fmt.Sscanf(s[0], "Tile %d:", &id)

		lines := strings.Split(s[1], "\n")
		size := len(lines)

		t := Tile{id: id, grid: make([][]rune, size), neighbors: make([]Tile, 0)}
		for i, line := range lines {
			t.grid[i] = make([]rune, size)
			for j, char := range line {
				t.grid[i][j] = char
			}
		}
		t.neighbors = findNeighbourTile(t, tileSet)

		tileSet[t.id] = t
	}

	return tileSet
}

func findNeighbourTile(tile Tile, tileSet TileSet) []Tile {
	neighbors := make([]Tile, 0)

	for _, otherTile := range tileSet {
		if intersect := utils.Intersect(tile.AllPossibleEdges(), otherTile.AllPossibleEdges()); len(intersect) > 0 {
			neighbors = append(neighbors, otherTile)

			tileSet[otherTile.id] = Tile{
				id:        otherTile.id,
				grid:      otherTile.grid,
				neighbors: append(tileSet[otherTile.id].neighbors, tile),
			}
		}
	}

	return neighbors
}
