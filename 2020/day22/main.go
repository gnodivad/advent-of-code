package main

import (
	"container/list"
	"fmt"
	"gnodivad/advent-of-code/utils"
	"strconv"
	"strings"
)

type Player struct {
	deck *list.List
}

type Combact struct {
	player1 *Player
	player2 *Player
}

func (p Player) String() string {
	var sb strings.Builder
	for e := p.deck.Front(); e != nil; e = e.Next() {
		sb.WriteString(strconv.Itoa(e.Value.(int)))
		if e.Next() != nil {
			sb.WriteString(",")
		}
	}

	return sb.String()
}

func (player Player) DrawTopCard() int {
	return player.deck.Remove(player.deck.Front()).(int)
}

func (player Player) CloneDeck(count int) *list.List {
	cloneDeck := list.New()
	for e := player.deck.Front(); e != nil && count != 0; e, count = e.Next(), count-1 {
		cloneDeck.PushBack(e.Value.(int))
	}

	return cloneDeck
}

func main() {
	fmt.Println("--- Part One ---")
	combact := initCombact(utils.ReadStringsFromFile("2020/day22/input.txt"))
	combact.run()

	fmt.Println("--- Part Two ---")
	combact = initCombact(utils.ReadStringsFromFile("2020/day22/input.txt"))
	winner := combact.runRecursively()
	announceWinnerScore(winner)
}

func (combact Combact) run() {
	for {
		player1Card, player2Card := combact.player1.DrawTopCard(), combact.player2.DrawTopCard()

		if player1Card > player2Card {
			combact.player1.deck.PushBack(player1Card)
			combact.player1.deck.PushBack(player2Card)
		} else {
			combact.player2.deck.PushBack(player2Card)
			combact.player2.deck.PushBack(player1Card)
		}

		if winner := combact.getWinner(); winner != nil {
			announceWinnerScore(winner)
			break
		}
	}
}

func (combact Combact) runRecursively() *Player {
	previousRounds := make(map[string]bool, combact.player1.deck.Len()+combact.player2.deck.Len())

	for {
		currentRoundString := combact.player1.String() + "|" + combact.player2.String()
		if _, exists := previousRounds[currentRoundString]; exists {
			return combact.player1
		}

		previousRounds[currentRoundString] = true

		player1Card, player2Card := combact.player1.DrawTopCard(), combact.player2.DrawTopCard()
		if player1Card <= combact.player1.deck.Len() && player2Card <= combact.player2.deck.Len() {
			newPlayer1, newPlayer2 := &Player{deck: combact.player1.CloneDeck(player1Card)}, &Player{deck: combact.player2.CloneDeck(player2Card)}
			subRound := Combact{player1: newPlayer1, player2: newPlayer2}

			winner := subRound.runRecursively()
			if winner == newPlayer1 {
				combact.player1.deck.PushBack(player1Card)
				combact.player1.deck.PushBack(player2Card)
			} else {
				combact.player2.deck.PushBack(player2Card)
				combact.player2.deck.PushBack(player1Card)
			}
		} else {
			if player1Card > player2Card {
				combact.player1.deck.PushBack(player1Card)
				combact.player1.deck.PushBack(player2Card)
			} else {
				combact.player2.deck.PushBack(player2Card)
				combact.player2.deck.PushBack(player1Card)
			}
		}

		if winner := combact.getWinner(); winner != nil {
			return winner
		}
	}
}

func (combact Combact) getWinner() *Player {
	if combact.player1.deck.Len() == 0 {
		return combact.player2
	} else if combact.player2.deck.Len() == 0 {
		return combact.player1
	}

	return nil
}

func announceWinnerScore(winner *Player) {
	score, scoreFactor := 0, winner.deck.Len()
	for e := winner.deck.Front(); e != nil; e, scoreFactor = e.Next(), scoreFactor-1 {
		score += (e.Value.(int) * scoreFactor)
	}

	fmt.Println(score)
}

func initCombact(inputs []string) Combact {
	combact := Combact{
		player1: &Player{deck: list.New()},
		player2: &Player{deck: list.New()},
	}

	inputZone := 0
	for _, input := range inputs {
		if input == "" {
			inputZone++
			continue
		}

		if strings.Contains(input, "Player ") {
			continue
		}

		if inputZone == 0 {
			combact.player1.deck.PushBack(utils.ParseInt(input))
		} else if inputZone == 1 {
			combact.player2.deck.PushBack(utils.ParseInt(input))
		}
	}

	return combact
}
