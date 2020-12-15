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

const size = 36

type Mask struct {
	BitArray [size]byte
	And      int
	Or       int
}

type Binary struct {
	n int
}

var regexForMemoryInstruction = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func main() {
	instructions := readStringsFromFile("day14/input.txt")

	memory1 := make(map[int]Binary)
	memory2 := make(map[int]Binary)

	var mask Mask
	for _, instructionAsString := range instructions {
		instruction := strings.Split(instructionAsString, " = ")
		operation, value := instruction[0], instruction[1]

		if operation == "mask" {
			mask = createNewMask(value)
		} else {
			p := regexForMemoryInstruction.FindStringSubmatch(instructionAsString)
			memoryAddress, _ := strconv.Atoi(p[1])
			n, _ := strconv.Atoi(value)

			b := Binary{n: n}
			memory1[memoryAddress] = mask.apply(b)

			possibleMemoryAddress := mask.applyOnMemoryAddress(memoryAddress)
			for _, m := range possibleMemoryAddress {
				memory2[m] = b
			}
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(calculateSum(memory1))

	fmt.Println("--- Part Two ---")
	fmt.Println(calculateSum(memory2))
}

func (mask Mask) apply(b Binary) Binary {
	bClone := b
	bClone.n &= mask.And
	bClone.n |= mask.Or

	return bClone
}

func (mask Mask) applyOnMemoryAddress(memoryAddress int) []int {
	memoryAddress |= mask.Or

	memoryAddressList := []int{memoryAddress}

	for i, bit := range mask.BitArray {
		if bit == 'X' {
			newMemoryAddressList := make([]int, 0, len(memoryAddressList)*2)

			for _, m := range memoryAddressList {
				// mask = 0X1, address = 011 and assume only 3 bit
				// 011 & 101 = 001 (1 << 3 = 1000 - 1 = 111 - 1 = 110 << 1 = 101)
				// 011 | 010 = 011 (1 << 1 = 10)
				newMemoryAddressList = append(newMemoryAddressList, m&((1<<size-1)-1<<(size-1-i)))
				newMemoryAddressList = append(newMemoryAddressList, m|1<<(size-1-i))
			}
			memoryAddressList = newMemoryAddressList
		}
	}

	return memoryAddressList
}

func calculateSum(memory map[int]Binary) int {
	sum := 0
	for _, binary := range memory {
		sum += binary.n
	}

	return sum
}

func createNewMask(maskAsString string) Mask {
	var mask Mask
	for i, char := range []byte(maskAsString) {
		mask.BitArray[i] = char

		if char == '1' {
			mask.Or += (1 << (size - 1 - i))
		}

		if char != '0' {
			mask.And += (1 << (size - 1 - i))
		}
	}

	return mask
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
