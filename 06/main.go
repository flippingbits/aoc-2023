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

func main() {
	input, err := os.Open(os.Getenv("AOC_INPUT_FILE"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	re := regexp.MustCompile("\\d+")

	scanner.Scan()
	time, _ := strconv.Atoi(strings.Join(re.FindAllString(scanner.Text(), -1)[:], ""))
	scanner.Scan()
	distance, _ := strconv.Atoi(strings.Join(re.FindAllString(scanner.Text(), -1)[:], ""))

	numberOfWays := 0
	for i := 1; i < time; i++ {
		travelTime := time - i
		travelDistance := travelTime * i
		if travelDistance > distance {
			numberOfWays++
		}
	}

	fmt.Println(numberOfWays)
}
