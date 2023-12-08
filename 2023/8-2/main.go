package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	chars, err := os.ReadFile("./2023/8-2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(chars), "\n")

	r1 := regexp.MustCompile(`(\w)`)
	r2 := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)

	instructions := r1.FindAllString(lines[0], -1)
	network := map[string][2]string{}
	nodes := []string{}
	for i := 2; i < len(lines); i++ {
		n := r2.FindAllStringSubmatch(lines[i], -1)
		network[n[0][1]] = [2]string{n[0][2], n[0][3]}

		if strings.HasSuffix(n[0][1], "A") {
			nodes = append(nodes, n[0][1])
		}
	}

	steps := 0
	i := 0
	nodesSteps := make([]int, len(nodes))
	nodesComplete := 0
	for {
		if i >= len(instructions) {
			i = 0
		}
		steps++

		for j, n := range nodes {
			if instructions[i] == "L" {
				nodes[j] = network[n][0]
			} else {
				nodes[j] = network[n][1]
			}
			if nodes[j][2:3] == "Z" {
				nodesSteps[j] = steps
				nodesComplete++
			}
		}

		if nodesComplete == len(nodesSteps) {
			break
		}

		i++
	}

	steps = 0

	slices.Sort(nodesSteps)

	result := 0
	for {
		divisor := 1
		for i := int(math.Min(float64(nodesSteps[0]), float64(nodesSteps[1]))); i > 0; i-- {
			if i == 1 || nodesSteps[0]%i == 0 && nodesSteps[1]%i == 0 {
				divisor = i
				break
			}
		}

		result = nodesSteps[0] * nodesSteps[1] / divisor

		if len(nodesSteps) <= 2 {
			fmt.Println("steps: ", result)
			break
		}

		nodesSteps = append([]int{result}, nodesSteps[2:]...)
	}
}
