package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {

	total := 0
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

	return "total: " + utils.ToString(total)
}

func part2(lines []string) string {
	total := 0
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

	return "total: " + utils.ToString(total)
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
