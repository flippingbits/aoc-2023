package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func isDigit(text byte) bool {
	r, _ := utf8.DecodeLastRuneInString(string(text))
	return unicode.IsDigit(r)
}

func getNumberFromPosition(lines []string, row int, column int) int {
	// find start of number
	for column > 0 && isDigit(lines[row][column-1]) {
		column--
	}

	// find end of number
	numberBuffer := ""
	for column < len(lines[row]) && isDigit(lines[row][column]) {
		numberBuffer += string(lines[row][column])
		column++
	}

	parsedNumber, _ := strconv.Atoi(numberBuffer)

	return parsedNumber
}

func getGearRatio(lines []string, row int, column int) int {
	var numbers []int

	// look to the left
	if column > 0 && isDigit(lines[row][column-1]) {
		numbers = append(numbers, getNumberFromPosition(lines, row, column-1))
	}

	// look to the right
	if column < len(lines[row])-1 && isDigit(lines[row][column+1]) {
		numbers = append(numbers, getNumberFromPosition(lines, row, column+1))
	}

	// look up
	if row > 0 {
		if column > 0 && isDigit(lines[row-1][column-1]) {
			numbers = append(numbers, getNumberFromPosition(lines, row-1, column-1))
		} else if isDigit(lines[row-1][column]) {
			numbers = append(numbers, getNumberFromPosition(lines, row-1, column))
		}
		if column < len(lines[row-1])-1 && !isDigit(lines[row-1][column]) && isDigit(lines[row-1][column+1]) {
			numbers = append(numbers, getNumberFromPosition(lines, row-1, column+1))
		}
	}

	// look down
	if row < len(lines)-1 {
		if column > 0 && isDigit(lines[row+1][column-1]) {
			numbers = append(numbers, getNumberFromPosition(lines, row+1, column-1))
		} else if isDigit(lines[row+1][column]) {
			numbers = append(numbers, getNumberFromPosition(lines, row+1, column))
		}
		if column < len(lines[row+1])-1 && !isDigit(lines[row+1][column]) && isDigit(lines[row+1][column+1]) {
			numbers = append(numbers, getNumberFromPosition(lines, row+1, column+1))
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}

	return 0
}

func main() {
	input, err := os.Open(os.Getenv("AOC_INPUT_FILE"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer input.Close()

	sum := 0
	var lines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for rowIndex, line := range lines {
		for columnIndex, char := range line {
			if char == '*' {
				sum += getGearRatio(lines, rowIndex, columnIndex)
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
