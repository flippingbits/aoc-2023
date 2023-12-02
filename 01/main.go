package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.Open(os.Getenv("AOC_INPUT_FILE"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	sum := 0
	textNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()

		numbersInLine := ""
		for lineIdx, char := range line {
			if unicode.IsDigit(char) {
				numbersInLine += string(char)
			} else {
				for idx, num := range textNumbers {
					if strings.HasPrefix(line[lineIdx:len(line)], num) {
						numbersInLine += strconv.Itoa(idx + 1)
						break
					}
				}
			}
		}
		// Convert buffer to number and add to sum
		number, err := strconv.Atoi(string(numbersInLine[0]) + string(numbersInLine[len(numbersInLine)-1]))
		if err == nil {
			sum += number
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
