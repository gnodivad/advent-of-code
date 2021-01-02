package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
)

var offsets = map[string]utils.VectorAxiaHexagonal{
	"e":  {Q: 1, R: 0},
	"se": {Q: 0, R: 1},
	"sw": {Q: -1, R: 1},
	"w":  {Q: -1, R: 0},
	"nw": {Q: 0, R: -1},
	"ne": {Q: 1, R: -1},
}

func main() {
	tiles := utils.ReadStringsFromFile("2020/day24/input.txt")

	blackTiles := getAllBlackTile(tiles)

	fmt.Println("--- Part One ---")
	fmt.Println(len(blackTiles))

	fmt.Println("--- Part Two ---")
	blackTiles = flipBlackTile(blackTiles, 100)
	fmt.Println(len(blackTiles))
}

func getAllBlackTile(tiles []string) map[utils.VectorAxiaHexagonal]bool {
	blackTile := make(map[utils.VectorAxiaHexagonal]bool)

	for _, tile := range tiles {
		tilePosition := parseTile(tile)

		if _, exist := blackTile[tilePosition]; !exist {
			blackTile[tilePosition] = true
		} else {
			delete(blackTile, tilePosition)
		}
	}

	return blackTile
}

func flipBlackTile(blackTiles map[utils.VectorAxiaHexagonal]bool, rounds int) map[utils.VectorAxiaHexagonal]bool {
	for rounds > 0 {
		neighborsTiles := make(map[utils.VectorAxiaHexagonal]int)

		for blackTile := range blackTiles {
			for _, offset := range offsets {
				neighborsTiles[blackTile.Add(offset)]++
			}
		}

		newBlackTiles := make(map[utils.VectorAxiaHexagonal]bool)
		for pos, neighborBlackTileCount := range neighborsTiles {
			if _, isBlack := blackTiles[pos]; (isBlack && neighborBlackTileCount == 1) || neighborBlackTileCount == 2 {
				newBlackTiles[pos] = true
			}
		}

		blackTiles = newBlackTiles
		rounds--
	}

	return blackTiles
}

func parseTile(tile string) (position utils.VectorAxiaHexagonal) {
	for i := 0; i < len(tile); i++ {

		direction := string(tile[i])
		if direction == "s" || direction == "n" {
			i += 1
			direction += string(tile[i])
		}

		position = position.Add(offsets[direction])
	}

	return
}
