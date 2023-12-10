package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {
	r1 := regexp.MustCompile(`(?m)(\-?\d+)`)

	total := 0
	for _, line := range lines {
		ss := r1.FindAllString(line, -1)

		nn := make([]int, len(ss))
		for i, n := range ss {
			nn[i], _ = strconv.Atoi(n)
		}
		nextN := madameIrmaPart1(nn)
		total += nextN
	}
	return utils.ToString(total)
}

func madameIrmaPart1(line []int) int {
	prevision := 0
	for {
		prevision += line[len(line)-1]
		line = calculateLinesDiffPart1(line)

		if slices.Max(line) == 0 && slices.Min(line) == 0 {
			break
		}
	}

	return prevision
}

func calculateLinesDiffPart1(line []int) []int {
	newLine := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		newLine[i-1] = line[i] - line[i-1]
	}
	return newLine
}

func part2(lines []string) string {
	r1 := regexp.MustCompile(`(?m)(\-?\d+)`)

	total := 0
	for _, line := range lines {
		ss := r1.FindAllString(line, -1)

		nn := make([]int, len(ss))
		for i, n := range ss {
			nn[i], _ = strconv.Atoi(n)
		}
		previousN := madameIrmaPart2(nn)
		total += previousN
	}
	return utils.ToString(total)
}

func madameIrmaPart2(line []int) int {
	allNewLines := [][]int{0: line}

	for {
		line = calculateLinesDiffPart2(line)
		allNewLines = append(allNewLines, line)

		if slices.Max(line) == 0 && slices.Min(line) == 0 {
			break
		}
	}

	slices.Reverse(allNewLines)
	for i, line := range allNewLines {
		if i == 0 {
			allNewLines[i] = append([]int{0}, line...)
		} else {
			allNewLines[i] = append([]int{line[0] - allNewLines[i-1][0]}, line...)
		}
	}
	return allNewLines[len(allNewLines)-1][0]
}

func calculateLinesDiffPart2(line []int) []int {
	newLine := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		newLine[i-1] = line[i] - line[i-1]
	}
	return newLine
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
