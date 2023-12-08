package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open(os.Getenv("AOC_INPUT_FILE"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer input.Close()

	sum := 0

	scanner := bufio.NewScanner(input)

	line := 0
	scratchCards := make(map[int]int)
	scratchCards[0] = 1

	for scanner.Scan() {
		card := strings.Split(strings.Split(scanner.Text(), ": ")[1], " | ")
		prefix := card[0]
		suffix := card[1]

		winningNumbers := make(map[int]bool)
		// process winning numbers of game
		for _, num := range strings.Split(prefix, " ") {
			if len(strings.TrimSpace(num)) > 0 {
				numInt, _ := strconv.Atoi(num)
				winningNumbers[numInt] = true
			}
		}

		// process numbers of card
		if _, ok := scratchCards[line]; !ok {
			scratchCards[line] = 1
		}
		copies := scratchCards[line]
		sum += copies
		for i := 0; i < copies; i++ {
			points := 0
			for _, num := range strings.Split(suffix, " ") {
				if len(strings.TrimSpace(num)) > 0 {
					numInt, _ := strconv.Atoi(num)
					// check if number is part of winningNumbers
					if _, ok := winningNumbers[numInt]; ok {
						points++
					}
				}
			}

			// determine number of copies for successive cards
			for j := 1; j <= points; j++ {
				scratchCard := line + j
				val, ok := scratchCards[scratchCard]
				if ok {
					scratchCards[scratchCard] = val + 1
				} else {
					scratchCards[scratchCard] = 2
				}
			}
		}

		line++
	}

	fmt.Printf("Sum: %d\n", sum)
}
