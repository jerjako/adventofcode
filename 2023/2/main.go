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
	chars, err := os.ReadFile("./2023/2/input.txt")
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
		game := strings.Split(l[0], " ")
		gameID, _ := strconv.Atoi(game[1])

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

		if totalColors["red"] <= 12 && totalColors["green"] <= 13 && totalColors["blue"] <= 14 {
			total += gameID
		}
	}

	fmt.Println("total: " + strconv.Itoa(total))
}
