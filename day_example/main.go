package main

import (
	"fmt"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {
	return "part1 result"
}

func part2(lines []string) string {
	return "part2 result"
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
