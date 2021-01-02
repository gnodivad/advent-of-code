package main

import (
	"gnodivad/advent-of-code/utils"
	"strings"
)

type TileSet map[int]Tile

type Tile struct {
	id        int
	grid      [][]rune
	neighbors []Tile
}

func (t Tile) Top() string {
	return string(t.grid[0])
}

func (t Tile) Bottom() string {
	return string(t.grid[len(t.grid)-1])
}

func (t Tile) Left() string {
	var left strings.Builder
	for i := 0; i < len(t.grid[0]); i++ {
		left.WriteRune(t.grid[i][0])
	}

	return left.String()
}

func (t *Tile) Right() string {
	var right strings.Builder
	size := len(t.grid[0])
	for i := 0; i < size; i++ {
		right.WriteRune(t.grid[i][size-1])
	}

	return right.String()
}

func (t Tile) Edges() []string {
	return []string{
		t.Top(),
		t.Bottom(),
		t.Left(),
		t.Right(),
	}
}

func (t Tile) AllPossibleEdges() []string {
	edges := t.Edges()

	return []string{
		edges[0],
		utils.ReverseString(edges[0]),
		edges[1],
		utils.ReverseString(edges[1]),
		edges[2],
		utils.ReverseString(edges[2]),
		edges[3],
		utils.ReverseString(edges[3]),
	}
}
