package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	chars, err := os.ReadFile("./2023/8/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(chars), "\n")

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

	fmt.Println("steps: " + strconv.Itoa(steps))
}
