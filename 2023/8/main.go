package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {

	r1 := regexp.MustCompile(`(\w)`)
	r2 := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

	instructions := r1.FindAllString(lines[0], -1)

	nodes := []string{}
	network := map[string][2]string{}
	node := ""
	for i := 2; i < len(lines); i++ {
		n := r2.FindAllStringSubmatch(lines[i], -1)
		network[n[0][1]] = [2]string{n[0][2], n[0][3]}
		if node == "" {
			node = n[0][1]
		}

		nodes = append(nodes, n[0][1])
	}

	var nodePtr *string
	var nodeZZZPtr *string
	networkPtr := map[*string][2]*string{}
	for n, networks := range network {
		for i := range nodes {
			if nodes[i] != n {
				continue
			}

			if n == "AAA" {
				nodePtr = &nodes[i]
			} else if n == "ZZZ" {
				nodeZZZPtr = &nodes[i]
			}

			var nL *string
			var nR *string
			for j := range nodes {
				if nodes[j] != networks[0] {
					continue
				}

				nL = &nodes[j]
				break
			}
			for j := range nodes {
				if nodes[j] != networks[1] {
					continue
				}

				nR = &nodes[j]
				break
			}
			networkPtr[&nodes[i]] = [2]*string{nL, nR}
		}

	}

	steps := 0
	i := 0
	for {
		if i >= len(instructions) {
			i = 0
		}

		if instructions[i] == "L" {
			nodePtr = networkPtr[nodePtr][0]
		} else {
			nodePtr = networkPtr[nodePtr][1]
		}

		steps++
		i++
		if nodePtr == nodeZZZPtr {
			break
		}
	}

	return utils.ToString(steps)
}

func part2(lines []string) string {

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
			return utils.ToString(result)
		}

		nodesSteps = append([]int{result}, nodesSteps[2:]...)
	}
}

func main() {
	lines, doPart1, doPart2 := utils.RunDay()
	if doPart1 {
		fmt.Println("result: ", part1(lines))
	}
	if doPart2 {
		fmt.Println("result: ", part2(lines))
	}
}
