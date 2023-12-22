package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

func getStrengthOfCard(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "T":
		return 10
	case "J":
		return 1
	default:
		strength, _ := strconv.Atoi(card)
		return strength
	}
}

func getStrength(hand Hand) int {
	freq := make(map[string]int)
	cards := strings.Split(hand.Cards, "")
	for _, card := range cards {
		freq[card] = freq[card] + 1
	}
	if len(freq) == 1 { // Five of a kind
		return 20
	} else if len(freq) == 2 { // Four of a kind or Full house
		for _, count := range freq {
			if count == 4 {
				// Four of a kind
				if freq["J"] >= 1 { // Make it a five of a kind
					return 20
				} else {
					return 19
				}
			}
		}
		// Full house
		if freq["J"] > 1 { // Makes it a five of a kind
			return 20
		} else {
			return 18
		}
	} else if len(freq) == 3 { // Three of a kind or two pair
		for _, count := range freq {
			if count == 3 {
				// Three of a kind
				if freq["J"] >= 1 { // Make it a four of  akind
					return 19
				} else {
					return 17
				}
			}
		}
		// Two pair
		if freq["J"] == 1 { // Make it a full house
			return 18
		} else if freq["J"] == 2 { // Make it a four of a kind
			return 19
		} else {
			return 16
		}
	} else if len(freq) == 4 { // One pair
		if freq["J"] >= 1 { // Make it a three of a kind
			return 17
		} else {
			return 15
		}
	} else { // High card
		if freq["J"] == 1 { // Make it a pair
			return 15
		} else {
			return 14
		}
	}
}

type ByStrength []Hand

func (a ByStrength) Len() int {
	return len(a)
}

func (a ByStrength) Less(i, j int) bool {
	iStrength := getStrength(a[i])
	jStrength := getStrength(a[j])
	if iStrength == jStrength { // secondary ordering
		iString := strings.Split(a[i].Cards, "")
		jString := strings.Split(a[j].Cards, "")
		for idx := 0; idx < len(iString); idx++ {
			if iString[idx] != jString[idx] {
				return getStrengthOfCard(iString[idx]) < getStrengthOfCard(jString[idx])
			}
		}
		return false
	} else {
		return iStrength < jStrength
	}
}

func (a ByStrength) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	input, err := os.Open(os.Getenv("AOC_INPUT_FILE"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer input.Close()

	var hands []Hand
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cards := line[0]
		bid, _ := strconv.Atoi(line[1])
		hands = append(hands, Hand{cards, bid})
	}
	sort.Sort(ByStrength(hands))

	total_winnings := 0
	for idx, hand := range hands {
		total_winnings += (idx + 1) * hand.Bid
	}

	fmt.Printf("Total winnings: %d\n", total_winnings)
}
