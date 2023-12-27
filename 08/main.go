package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(numbers []int) int {
	res := numbers[0] * numbers[1] / gcd(numbers[0], numbers[1])
	for i := 2; i < len(numbers); i++ {
		res = res * numbers[i] / gcd(res, numbers[i])
	}
	return res
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
	instructions := scanner.Text()
	scanner.Scan()

	nodes := make(map[string]Node)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		directions := strings.Split(line[1][1:len(line[1])-1], ", ")
		nodes[line[0]] = Node{
			Left:  directions[0],
			Right: directions[1],
		}
	}

	steps := 0
	var currentNodes []string
	// determine starting points
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			currentNodes = append(currentNodes, node)
		}
	}

	var stepsPerNode []int
	for {
		keepSearching := false
		for _, ch := range instructions {
			// move all nodes into the given direction
			for idx, node := range currentNodes {
				if len(node) > 0 {
					if string(ch) == "L" {
						currentNodes[idx] = nodes[node].Left
					} else {
						currentNodes[idx] = nodes[node].Right
					}
				}
			}

			steps++

			// check if we reached the destination for any of the paths
			for idx, node := range currentNodes {
				if len(node) > 0 && strings.HasSuffix(node, "Z") {
					currentNodes[idx] = ""
					stepsPerNode = append(stepsPerNode, steps)
				}
			}

			// check if we need to keep searching for a destination
			for _, node := range currentNodes {
				if len(node) > 0 {
					keepSearching = true
					break
				}
			}
		}

		if !keepSearching {
			break
		}
	}

	// Determine least common multiple
	fmt.Println(lcm(stepsPerNode))
}
