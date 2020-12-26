package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
)

func main() {
	input := utils.ReadStringsFromFile("2020/day17/input.txt")

	fmt.Println("--- Part One ---")
	worldIn3D := initWorldWithDimension(input, 3)
	worldIn3D.run(6)

	fmt.Println("--- Part Two ---")
	worldIn4D := initWorldWithDimension(input, 4)
	worldIn4D.run(6)
}

type World struct {
	locations map[utils.Vector]bool
}

func initWorldWithDimension(input []string, dimension int) World {
	var world World
	world.locations = make(map[utils.Vector]bool)

	for y, line := range input {
		for x, char := range line {
			if char != '#' {
				continue
			}

			if dimension == 3 {
				world.locations[utils.Vector3{X: x, Y: y, Z: 0}] = true
			} else if dimension == 4 {
				world.locations[utils.Vector4{X: x, Y: y, Z: 0, W: 0}] = true
			}
		}
	}

	return world
}

func (world World) run(times int) {
	for i := 0; i < times; i++ {
		world.simulate()
	}
	fmt.Println(world.calculateActiveLocation())
}

func (world *World) simulate() {
	lookup := make(map[utils.Vector]bool)
	for vector, active := range world.locations {
		lookup[vector] = active

		for _, neighbor := range vector.Neighbors() {
			if _, ok := lookup[neighbor]; !ok {
				lookup[neighbor] = false
			}
		}
	}

	newLocations := make(map[utils.Vector]bool)

	for vector, active := range lookup {
		activeNeighbors := 0
		for _, neighbor := range vector.Neighbors() {
			if lookup[neighbor] {
				activeNeighbors++
			}

		}
		if active && (activeNeighbors == 2 || activeNeighbors == 3) {
			newLocations[vector] = true
		} else if !active && activeNeighbors == 3 {
			newLocations[vector] = true
		}
	}

	world.locations = newLocations
}

func (world World) calculateActiveLocation() int {
	count := 0
	for _, active := range world.locations {
		if active {
			count++
		}
	}

	return count
}
