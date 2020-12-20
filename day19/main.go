package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
	"strconv"
	"strings"
)

type Rule struct {
	subRule1, subRule2 []int
	char               rune
}

func (rule Rule) String() string {
	return fmt.Sprintln(rule.subRule1, rule.subRule2, string(rule.char))
}

func main() {
	inputs := utils.ReadStringsFromFile("day19/input.txt")

	zone := 0

	rules := make(map[int]Rule)
	messages := make([]string, 0)

	for _, input := range inputs {
		if input == "" {
			zone++
			continue
		}

		if zone == 0 {
			inputSlice := strings.Split(input, ": ")
			index, _ := strconv.Atoi(inputSlice[0])
			rules[index] = parseRawRule(inputSlice[1])
		} else if zone == 1 {
			messages = append(messages, input)
		}
	}

	fmt.Println("--- Part One ---")
	validateAll(messages, rules)

	fmt.Println("--- Part Two ---")
	rules[8] = Rule{subRule1: []int{42}, subRule2: []int{42, 8}}
	rules[11] = Rule{subRule1: []int{42, 31}, subRule2: []int{42, 11, 31}}
	validateAll(messages, rules)
}

func validateAll(messages []string, rules map[int]Rule) {
	correctCount := 0
	for _, message := range messages {
		if validate([]rune(message), 0, rules, 0, []int{}) {
			correctCount++
		}
	}

	fmt.Println(correctCount)
}

func validate(message []rune, charIndex int, rules map[int]Rule, ruleIndex int, remainingRule []int) bool {
	rule := rules[ruleIndex]
	if rule.char != 0 {
		if message[charIndex] != rule.char {
			return false
		}

		if len(remainingRule) == 0 {
			return charIndex == len(message)-1
		} else if charIndex+1 >= len(message) {
			return false
		} else {
			return validate(message, charIndex+1, rules, remainingRule[0], remainingRule[1:])
		}
	}

	toAdd := rule.subRule1[1:]
	newRemainingRule := make([]int, len(toAdd)+len(remainingRule))
	copy(newRemainingRule, toAdd)
	copy(newRemainingRule[len(toAdd):], remainingRule)
	if validate(message, charIndex, rules, rule.subRule1[0], newRemainingRule) {
		return true
	}

	if len(rule.subRule2) != 0 {
		toAdd := rule.subRule2[1:]
		newRemainingRule := make([]int, len(toAdd)+len(remainingRule))
		copy(newRemainingRule, toAdd)
		copy(newRemainingRule[len(toAdd):], remainingRule)

		if validate(message, charIndex, rules, rule.subRule2[0], newRemainingRule) {
			return true
		}
	}

	return false
}

func parseRawRule(rawRule string) (rule Rule) {
	if strings.Contains(rawRule, `"`) {
		rule.char = rune(rawRule[1])

		return
	}

	if strings.Contains(rawRule, ` | `) {
		for i, subRuleList := range strings.Split(rawRule, ` | `) {
			if i == 0 {
				rule.subRule1 = parseRawRuleInList(subRuleList)
			} else if i == 1 {
				rule.subRule2 = parseRawRuleInList(subRuleList)
			}
		}

		return
	}

	rule.subRule1 = parseRawRuleInList(rawRule)

	return
}

func parseRawRuleInList(rawRule string) (subRule []int) {
	subRule = make([]int, 0)

	for _, subRuleAsString := range strings.Split(rawRule, ` `) {
		subRuleAsInt, _ := strconv.Atoi(subRuleAsString)
		subRule = append(subRule, subRuleAsInt)
	}

	return
}
