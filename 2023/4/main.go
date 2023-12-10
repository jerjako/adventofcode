package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {
	total := 0
	for _, line := range lines {
		line = strings.Replace(line, "  ", " ", -1)
		l := strings.Split(line, " | ")
		l[0] = strings.Split(l[0], ": ")[1]

		winnings := strings.Split(l[0], " ")
		for i := range winnings {
			winnings[i] = strings.TrimSpace(winnings[i])
		}

		have := strings.Split(l[1], " ")

		winningCnt := 0
		for _, me := range have {
			if slices.Contains(winnings, strings.TrimSpace(me)) {
				if winningCnt == 0 {
					winningCnt = 1
				} else {
					winningCnt *= 2
				}
			}
		}
		total += winningCnt
	}

	return "total: " + utils.ToString(total)
}

func part2(lines []string) string {
	total := 0
	copies := make([]int, len(lines))
	for index, line := range lines {
		line = strings.Replace(line, "  ", " ", -1)
		l := strings.Split(line, " | ")
		l[0] = strings.Split(l[0], ": ")[1]

		winnings := strings.Split(l[0], " ")
		for i := range winnings {
			winnings[i] = strings.TrimSpace(winnings[i])
		}

		have := strings.Split(l[1], " ")

		winningCnt := 0
		for _, me := range have {
			if slices.Contains(winnings, strings.TrimSpace(me)) {
				winningCnt++
			}
		}
		for i := 0; i <= winningCnt; i++ {
			if i == 0 {
				copies[index+i]++
			} else {
				copies[index+i] += copies[index]
			}
		}
		total += copies[index]
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
