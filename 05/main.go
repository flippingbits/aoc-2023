package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
}

func applyMappings(mappings []Mapping, intervals [][]int) [][]int {
	newIntervals := make([][]int, 0)
	for idx, interval := range intervals {
		for _, mapping := range mappings {
			if interval[0] >= mapping.sourceStart && interval[1] <= mapping.sourceEnd {
				newStart := mapping.destinationStart + interval[0] - mapping.sourceStart
				newEnd := newStart + interval[1] - interval[0]
				intervals[idx] = []int{newStart, newEnd}
			} else if interval[0] >= mapping.sourceStart && interval[0] <= mapping.sourceEnd && interval[1] >= mapping.sourceEnd {
				newStart := mapping.destinationStart + interval[0] - mapping.sourceStart
				newEnd := mapping.destinationEnd
				intervals[idx] = []int{newStart, newEnd}
				newIntervals = append(
					newIntervals,
					[]int{mapping.sourceEnd + 1, interval[1]})
			} else if interval[0] < mapping.sourceStart && interval[1] >= mapping.sourceStart && interval[1] <= mapping.sourceEnd {
				newStart := mapping.destinationStart
				newEnd := mapping.destinationStart + interval[1] - mapping.sourceStart
				intervals[idx] = []int{newStart, newEnd}
				newIntervals = append(
					newIntervals,
					[]int{interval[0], mapping.sourceStart - 1})
			} else if interval[0] < mapping.sourceStart && interval[1] >= mapping.sourceEnd {
				intervals[idx] = []int{mapping.destinationStart, mapping.destinationEnd}
				newIntervals = append(
					newIntervals,
					[]int{interval[0], mapping.sourceStart - 1},
					[]int{mapping.sourceEnd + 1, interval[1]})
			}
		}
	}

	if len(newIntervals) > 0 {
		newIntervals = applyMappings(mappings, newIntervals)
		intervals = append(intervals, newIntervals...)
	}

	return intervals
}

func main() {
	input, err := os.Open(os.Getenv("AOC_INPUT_FILE"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	seedNumPairs := strings.Split(scanner.Text()[7:], " ")
	var intervals [][]int
	for idx, pair := range seedNumPairs {
		if idx%2 == 0 {
			seed, _ := strconv.Atoi(pair)
			num, _ := strconv.Atoi(seedNumPairs[idx+1])
			intervals = append(intervals, []int{seed, seed + num - 1})
		}
	}

	var mappings []Mapping
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if strings.HasSuffix(line, " map:") {
			intervals = applyMappings(mappings, intervals)
			mappings = make([]Mapping, 0)
			continue
		}

		mapping := strings.Split(line, " ")
		destinationStart, _ := strconv.Atoi(mapping[0])
		sourceStart, _ := strconv.Atoi(mapping[1])
		delta, _ := strconv.Atoi(mapping[2])

		mappings = append(
			mappings,
			Mapping{
				sourceStart:      sourceStart,
				sourceEnd:        sourceStart + delta - 1,
				destinationStart: destinationStart,
				destinationEnd:   destinationStart + delta - 1})
	}

	intervals = applyMappings(mappings, intervals)

	min := intervals[0][0]
	for _, interval := range intervals[1:] {
		if interval[0] < min {
			min = interval[0]
		}
	}
	fmt.Println(min)
}
