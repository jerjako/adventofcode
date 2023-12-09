package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	v  = true
	vv = true
)

func main() {
	chars, err := os.ReadFile("./2023/9/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(chars), "\n")

	r1 := regexp.MustCompile(`(?m)(\-?\d+)`)

	total := 0
	for _, line := range lines {
		if vv {
			fmt.Println(line)
		}
		ss := r1.FindAllString(line, -1)

		nn := make([]int, len(ss))
		for i, n := range ss {
			nn[i], _ = strconv.Atoi(n)
		}
		nextN := madameIrma(nn)
		if v {
			fmt.Println(line, " >> ", nextN)
		}
		total += nextN
	}
	fmt.Println("Total: ", total)
}

func madameIrma(line []int) int {
	prevision := 0
	for {
		prevision += line[len(line)-1]
		if vv {
			fmt.Println("before: ", line)
		}
		line = calculateLinesDiff(line)
		if vv {
			fmt.Println("after: ", line)
		}

		if slices.Max(line) == 0 && slices.Min(line) == 0 {
			break
		}
	}

	return prevision
}

func calculateLinesDiff(line []int) []int {
	newLine := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		newLine[i-1] = line[i] - line[i-1]
	}
	return newLine
}

// 2148393188 too high
// 1789635360 too high
// 653688137 too low
// 1789635132
