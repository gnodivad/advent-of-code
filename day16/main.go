package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
	"regexp"
	"strconv"
	"strings"
)

type Puzzle struct {
	rules        []*Rule // Store the address of rule object to make update pos easier in Part 2
	myTicket     Ticket
	otherTickets []Ticket
}

type Ticket []int

type Rule struct {
	name       string
	pos        int
	boundaries [2][2]int
}

func main() {
	puzzle := createPuzzle(utils.ReadStringsFromFile("day16/input.txt"))

	fmt.Println("--- Part One ---")
	fmt.Println(puzzle.calculateErrorRate())

	fmt.Println("--- Part Two ---")
	fmt.Println(puzzle.getMultiplicationOfAllDepartureField())
}

func createPuzzle(notes []string) Puzzle {
	p := Puzzle{}

	zone := 0

	for _, note := range notes {
		if note == "" {
			zone++
			continue
		}

		if strings.Contains(note, "your ticket:") || strings.Contains(note, "nearby tickets:") {
			continue
		}

		if zone == 0 {
			p.addNewRule(note)
		} else if zone == 1 {
			p.myTicket = createNewTicket(note)
		} else if zone == 2 {
			p.otherTickets = append(p.otherTickets, createNewTicket(note))
		}
	}

	return p
}

var regexForField = regexp.MustCompile(`(\D+):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)`)

func (p *Puzzle) addNewRule(note string) {
	r := regexForField.FindStringSubmatch(note)
	g1Lower, _ := strconv.Atoi(r[2])
	g1Upper, _ := strconv.Atoi(r[3])
	g2Lower, _ := strconv.Atoi(r[4])
	g2Upper, _ := strconv.Atoi(r[5])

	p.rules = append(p.rules, &Rule{
		name:       r[1],
		pos:        -1,
		boundaries: [2][2]int{{g1Lower, g1Upper}, {g2Lower, g2Upper}},
	})
}

func createNewTicket(note string) Ticket {
	t := make(Ticket, 0)

	for _, v := range strings.Split(note, ",") {
		t = append(t, utils.ParseInt(v))
	}

	return t
}

func (p Puzzle) calculateErrorRate() int {
	errorRate := 0
	for _, ticket := range p.otherTickets {
		if isValid, invalidValue := ticket.validateByRules(p.rules); !isValid {
			errorRate += invalidValue
		}
	}

	return errorRate
}

func (r Rule) validate(val int) bool {
	return (val >= r.boundaries[0][0] && val <= r.boundaries[0][1]) || val >= r.boundaries[1][0] && val <= r.boundaries[1][1]
}

func (ticket Ticket) validateByRules(rules []*Rule) (bool, int) {
	for _, t := range ticket {
		isValid := false

		for _, r := range rules {
			if r.validate(t) {
				isValid = true
			}
		}

		if !isValid {
			return false, t
		}
	}

	return true, 0
}

func (puzzle *Puzzle) getUnidentifiedRules() map[*Rule]bool {
	ret := make(map[*Rule]bool)

	for _, rule := range puzzle.rules {
		if rule.pos == -1 {
			ret[rule] = true
		}
	}

	return ret
}

func (puzzle Puzzle) identifyFields() {
	validTickets := make([]Ticket, 0)
	for _, ticket := range puzzle.otherTickets {
		if isValid, _ := ticket.validateByRules(puzzle.rules); isValid {
			validTickets = append(validTickets, ticket)
		}
	}

	possibleRules := puzzle.getUnidentifiedRules()

	for len(possibleRules) > 0 {
		for pos := 0; pos < len(puzzle.myTicket); pos++ {
			// For each of the rule, go through all the nearby tickets in specific position
			for _, rule := range puzzle.rules {
				for _, ticket := range validTickets {
					// If the rule does not match, exclude it from possible rules list and continue to next rule
					if !rule.validate(ticket[pos]) {
						delete(possibleRules, rule)
						break
					}
				}
			}

			if len(possibleRules) == 1 {
				for rule := range possibleRules {
					rule.pos = pos
				}
			}

			possibleRules = puzzle.getUnidentifiedRules()
		}
	}
}

func (puzzle Puzzle) getMultiplicationOfAllDepartureField() int {
	puzzle.identifyFields()

	result := 1

	for _, rule := range puzzle.rules {
		if strings.Contains(rule.name, "departure") {
			result *= puzzle.myTicket[rule.pos]
		}
	}

	return result
}
