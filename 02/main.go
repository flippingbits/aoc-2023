package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
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

	scanner := bufio.NewScanner(input)
	sum := float64(0)

	cubeRe := regexp.MustCompile(`(\d+) (\S+)`)
	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, ": ")
		suffix := lineParts[1]

		games := strings.Split(suffix, "; ")

		// keep track of max cubes needed for each color
		var maxRedCubes float64 = 0
		var maxGreenCubes float64 = 0
		var maxBlueCubes float64 = 0

		// parse games
		for _, game := range games {
			cubes := strings.Split(game, ", ")
			for _, cube := range cubes {
				pick := cubeRe.FindStringSubmatch(cube)
				count, _ := strconv.ParseFloat(pick[1], 64)
				color := pick[2]
				if color == "red" {
					maxRedCubes = math.Max(maxRedCubes, count)
				} else if color == "green" {
					maxGreenCubes = math.Max(maxGreenCubes, count)
				} else if color == "blue" {
					maxBlueCubes = math.Max(maxBlueCubes, count)
				}
			}
		}

		sum += (maxRedCubes * maxGreenCubes * maxBlueCubes)
	}

	fmt.Printf("Sum: %d\n", sum)
}
