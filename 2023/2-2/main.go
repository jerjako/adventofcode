package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	newLine = byte(10)
)

func main() {
	chars, err := os.ReadFile("./2023/2-2/input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	lines := strings.Split(string(chars), "\n")

	for _, line := range lines {
		totalColors := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		l := strings.Split(line, ": ")

		subsets := strings.Split(l[1], "; ")
		for _, subset := range subsets {
			cubes := strings.Split(subset, ", ")

			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				cubeCount, _ := strconv.Atoi(c[0])
				cubeColor := c[1]

				if cubeCount > totalColors[cubeColor] {
					totalColors[cubeColor] = cubeCount
				}
			}
		}

		total += totalColors["red"] * totalColors["green"] * totalColors["blue"]
	}

	fmt.Println("total: " + strconv.Itoa(total))
}
