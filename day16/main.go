package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Boundary struct {
	low  int
	high int
}

type Rule struct {
	name       string
	boundaries []Boundary
}

var regexForField = regexp.MustCompile(`(\D+):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)`)

func main() {
	notes := readStringsFromFile("day16/input.txt")

	rules := make([]Rule, 0)
	myTicket := make([]int, 0)

	isField, isTicket, isNearbyTicket := true, false, false
	possibleRules := make(map[string][]bool)

	sum := 0
	for _, note := range notes {
		if note == "your ticket:" {
			isField = false
			isTicket = true
		}

		if note == "nearby tickets:" {
			isTicket = false
			isNearbyTicket = true
		}

		if note == "" || note == "your ticket:" || note == "nearby tickets:" {
			continue
		}

		if isField {
			rule := createNewRule(note)
			rules = append(rules, rule)
			possibleRules[rule.name] = []bool{}
		}

		if isTicket {
			for _, ticketAsString := range strings.Split(note, ",") {
				ticketAsInt, _ := strconv.Atoi(ticketAsString)
				myTicket = append(myTicket, ticketAsInt)
			}

			for name := range possibleRules {
				possibleRules[name] = make([]bool, len(myTicket))
				for i := range possibleRules[name] {
					possibleRules[name][i] = true
				}
			}
		}

		if isNearbyTicket {
			tickets := make([]int, 0)
			var notPossibleRulesline [][]string
			var valid bool
			for _, ticketAsString := range strings.Split(note, ",") {
				ticketAsInt, _ := strconv.Atoi(ticketAsString)
				tickets = append(tickets, ticketAsInt)

				valid = false
				var notPossibleRules []string
				for _, rule := range rules {
					if rule.isFullfilledBy(ticketAsInt) {
						valid = true
					} else {
						notPossibleRules = append(notPossibleRules, rule.name)
					}
				}
				notPossibleRulesline = append(notPossibleRulesline, notPossibleRules)
			}

			if valid {
				for i, notPossibleRules := range notPossibleRulesline {
					for _, name := range notPossibleRules {
						possibleRules[name][i] = false
					}
				}
			}

			sum += getInvalidField(tickets, rules)
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(sum)

	fmt.Println("--- Part Two ---")

	fmt.Println(possibleRules)
}

func createNewRule(note string) Rule {
	r := regexForField.FindStringSubmatch(note)
	g1Lower, _ := strconv.Atoi(r[2])
	g1Upper, _ := strconv.Atoi(r[3])
	g2Lower, _ := strconv.Atoi(r[4])
	g2Upper, _ := strconv.Atoi(r[5])

	return Rule{
		name:       r[1],
		boundaries: []Boundary{{low: g1Lower, high: g1Upper}, {low: g2Lower, high: g2Upper}},
	}
}

func assignFields(nearbyTickets [][]int, rules []Rule) {
	allPossibleField := getAllPossibleFields(0, nearbyTickets[0], rules, make([][]string, 0), make([]string, 0), make(map[string]bool))

	fmt.Println(allPossibleField)
}

func getAllPossibleFields(index int, ticket []int, rules []Rule, allPossibleField [][]string, currentField []string, lookup map[string]bool) [][]string {
	if index == len(ticket) {
		allPossibleField = append(allPossibleField, currentField)
		return allPossibleField
	}

	for _, rule := range rules {
		if _, ok := lookup[rule.name]; !ok {
			if rule.isFullfilledBy(ticket[index]) {
				lookup[rule.name] = true
				cloneField := make([]string, len(currentField))
				copy(cloneField, currentField)
				cloneField = append(cloneField, rule.name)

				allPossibleField = getAllPossibleFields(index+1, ticket, rules, allPossibleField, cloneField, lookup)
			}
			delete(lookup, rule.name)
		}
	}

	return allPossibleField
}

func getInvalidField(tickets []int, rules []Rule) int {
	for _, ticketAsInt := range tickets {
		isValidField := false

		for _, rule := range rules {
			if rule.isFullfilledBy(ticketAsInt) {
				isValidField = true
				break
			}
		}
		if !isValidField {
			return ticketAsInt
		}
	}

	return 0
}

func (rule Rule) isFullfilledBy(value int) bool {
	for _, boundary := range rule.boundaries {
		if boundary.validate(value) {
			return true
		}
	}
	return false
}

func (boundary Boundary) validate(value int) bool {
	return value >= boundary.low && value <= boundary.high
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
