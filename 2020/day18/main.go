package main

import (
	"container/list"
	"fmt"
	"gnodivad/advent-of-code/utils"
	"strconv"
)

func main() {
	inputs := utils.ReadStringsFromFile("2020/day18/input.txt")

	fmt.Println("--- Part One ---")
	fmt.Println(calculateWhenAddAndMultiplyHaveSamePrecendence(inputs))

	fmt.Println("--- Part One ---")
	fmt.Println(calculateWhenAddHaveHigherPrecendence(inputs))
}

func calculateWhenAddAndMultiplyHaveSamePrecendence(inputs []string) int {
	return calculate(inputs, 1)
}

func calculateWhenAddHaveHigherPrecendence(inputs []string) int {
	return calculate(inputs, 2)
}

func calculate(inputs []string, mode int) int {
	sum := 0
	for _, input := range inputs {
		operands := list.New()
		operators := list.New()

		for _, char := range input {
			if char == ' ' {
				continue
			}

			if char == ')' {
				tempOperands := list.New()
				tempOperators := list.New()

				for operators.Back().Value != '(' {
					tempOperators.PushFront(operators.Remove(operators.Back()))

					if tempOperands.Len() == 0 {
						tempOperands.PushFront(operands.Remove(operands.Back()))
					}
					tempOperands.PushFront(operands.Remove(operands.Back()))
				}

				var result int
				if mode == 1 {
					result = computeResultInMemory(tempOperands, tempOperators)
				} else {
					result = computeResultInMemoryWithAdditionMorePrecedence(tempOperands, tempOperators)
				}

				operands.PushBack(result)
				// delete the first matching '('
				operators.Remove(operators.Back())
			} else {
				number, err := strconv.Atoi(string(char))
				if err != nil {
					operators.PushBack(char)
				} else {
					operands.PushBack(number)
				}
			}
		}

		if mode == 1 {
			sum += computeResultInMemory(operands, operators)
		} else {
			sum += computeResultInMemoryWithAdditionMorePrecedence(operands, operators)
		}
	}

	return sum
}

func computeResultInMemory(operands *list.List, operators *list.List) int {
	for operators.Front() != nil {
		operator := operators.Remove(operators.Front()).(rune)
		operand1 := operands.Remove(operands.Front()).(int)
		operand2 := operands.Remove(operands.Front()).(int)

		result := compute(operand1, operand2, operator)
		operands.PushFront(result)
	}

	return operands.Front().Value.(int)
}

func computeResultInMemoryWithAdditionMorePrecedence(operands *list.List, operators *list.List) int {
	operatorItr := operators.Front()
	operandItr := operands.Front()

	for operatorItr != nil {
		if operatorItr.Value.(rune) == '+' {
			temp := operatorItr.Next()
			operator := operators.Remove(operatorItr).(rune)
			operatorItr = temp

			temp = operandItr.Next().Next()
			operand2 := operands.Remove(operandItr.Next()).(int)
			operand1 := operands.Remove(operandItr).(int)

			result := compute(operand1, operand2, operator)
			if temp != nil {
				operandItr = operands.InsertBefore(result, temp)
			} else {
				operandItr = operands.PushBack(result)
			}
		} else {
			operandItr = operandItr.Next()
			operatorItr = operatorItr.Next()
		}
	}

	return computeResultInMemory(operands, operators)
}

func compute(operand1 int, operand2 int, operator rune) int {
	if operator == '+' {
		return operand1 + operand2
	}

	return operand1 * operand2
}
