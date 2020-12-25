package main

import (
	"container/list"
	"fmt"
	"gnodivad/advent-of-code/utils"
	"strconv"
	"strings"
)

type Game struct {
	cups     *list.List
	current  *list.Element
	position []*list.Element
}

func main() {
	fmt.Println("--- Part One ---")
	game := initGame(utils.ReadStringsFromFile("day23/input.txt")[0])
	game.play(100)
	fmt.Println(game.getResult())

	fmt.Println("--- Part Two ---")
	game = initGameWithExtraCups(utils.ReadStringsFromFile("day23/input.txt")[0], 1000000)
	game.play(10000000)
	fmt.Println(game.findTheStar())
}

func (game Game) play(round int) {
	for round != 0 {
		pickupCups := game.pickCupInClockwise(3)
		destinationCup := game.getDestinationCup(pickupCups)

		cupAfterInsert := destinationCup
		for _, cup := range pickupCups {
			game.position[cup] = game.cups.InsertAfter(cup, cupAfterInsert)
			cupAfterInsert = cupAfterInsert.Next()
		}

		game.current = game.nextCupInClockwise(nil)
		round--
	}
}

func (game Game) getResult() string {
	var sb strings.Builder
	p := game.position[1].Next()
	for {
		if p == game.position[1] {
			break
		}

		sb.WriteString(strconv.Itoa(p.Value.(int)))
		p = game.nextCupInClockwise(p)
	}

	return sb.String()
}

func (game Game) findTheStar() int {
	nextCup := game.position[1].Next()
	cupAfterNext := game.nextCupInClockwise(nextCup)

	return nextCup.Value.(int) * cupAfterNext.Value.(int)
}

func (game Game) getDestinationCup(excludes []int) *list.Element {
	current := game.current.Value.(int)
	for {
		current--

		if current == 0 {
			current = len(game.position) - 1
		}

		if !utils.Contains(excludes, current) {
			return game.position[current]
		}
	}
}

func (game Game) pickCupInClockwise(count int) []int {
	cups := make([]int, count)

	for i := 0; i < count; i++ {
		next := game.nextCupInClockwise(nil)
		cups[i] = next.Value.(int)
		game.cups.Remove(next)
	}

	return cups
}

func (game Game) nextCupInClockwise(current *list.Element) *list.Element {
	if current == nil {
		current = game.current
	}

	if current.Next() != nil {
		return current.Next()
	}

	return game.cups.Front()
}

func initGame(input string) Game {
	return initGameWithExtraCups(input, len(input))
}

func initGameWithExtraCups(input string, cupQuantity int) Game {
	if cupQuantity == 0 {
		cupQuantity = len(input)
	}

	cups := list.New()
	position := make([]*list.Element, cupQuantity+1)

	for _, char := range input {
		cupLabel := char - '0'
		position[cupLabel] = cups.PushBack(int(cupLabel))
	}

	for i := len(input) + 1; i <= cupQuantity; i++ {
		position[i] = cups.PushBack(i)
	}

	return Game{
		cups:     cups,
		current:  cups.Front(),
		position: position,
	}
}
