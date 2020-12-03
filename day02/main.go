package main

import (
	"strconv"
	"log"
    "bufio"
    "fmt"
    "strings"
	"os"
)

func main() {
    passwordDatabase := readStringsFromFile("day02/input.txt")
    
    validPasswordForMinMaxLimitPolicy := 0
    validPasswordForCharDefineInAnyPostionPolicy := 0
    for _, passwordRow := range passwordDatabase {
        passwordRowSlice := strings.Split(passwordRow, ": ")
        policy := passwordRowSlice[0]
        password := passwordRowSlice[1]

        policySlice := strings.Split(policy, " ")
        charLimit := policySlice[0]
        char := policySlice[1]

        charLimitSlice := strings.Split(charLimit, "-")
        min, _ := strconv.Atoi(charLimitSlice[0])
        max, _ := strconv.Atoi(charLimitSlice[1])

        if (isPasswordWithinMinAndMaxLimit(password, char, min, max)) {
            validPasswordForMinMaxLimitPolicy ++
        }

        if (isCharDefineInAnyPostion(password, char, min, max)) {
            validPasswordForCharDefineInAnyPostionPolicy ++
        }
    }

    fmt.Println("--- Part One ---")
    fmt.Println(validPasswordForMinMaxLimitPolicy)

    fmt.Println("--- Part Two ---")
    fmt.Println(validPasswordForCharDefineInAnyPostionPolicy)
}

func isPasswordWithinMinAndMaxLimit(password string, char string, min int, max int) (isCorrect bool) {
    isCorrect = false

    count := 0
    for _, c := range password {
        if (string(c) == char) {
            count++
        }
    }

    if (count >= min && count <= max) {
        isCorrect = true
    }

    return
}

func isCharDefineInAnyPostion(password string, char string, posA int, posB int) (isCorrect bool) {
    isCorrect = false

    if (string(password[posA - 1]) == char) {
        isCorrect = !isCorrect
    }

    if (string(password[posB - 1]) == char) {
        isCorrect = !isCorrect
    }

    return
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
