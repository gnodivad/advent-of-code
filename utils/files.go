package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadStringsFromFile(filepath string) (strings []string) {
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

func ReadFile(filepath string) string {
	buf, _ := ioutil.ReadFile(filepath)

	return strings.TrimSpace(string(buf))
}
