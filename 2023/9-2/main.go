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
	v  = false
	vv = false
)

func main() {
	chars, err := os.ReadFile("./2023/9-2/input.txt")
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
		previousN := madameIrma(nn)
		if v {
			fmt.Println(previousN, " >> ", line)
		}
		total += previousN
	}
	fmt.Println("Total: ", total)
}

func madameIrma(line []int) int {
	allNewLines := [][]int{0: line}

	for {
		line = calculateLinesDiff(line)
		allNewLines = append(allNewLines, line)
		if vv {
			fmt.Println("after: ", line)
		}

		if slices.Max(line) == 0 && slices.Min(line) == 0 {
			break
		}
	}

	fmt.Println(allNewLines)

	slices.Reverse(allNewLines)
	fmt.Println("start")
	for i, line := range allNewLines {
		if i == 0 {
			fmt.Println("add ZERO")
			allNewLines[i] = append([]int{0}, line...)
		} else {
			fmt.Println("add ", allNewLines[i-1][0]-line[0])
			allNewLines[i] = append([]int{line[0] - allNewLines[i-1][0]}, line...)
		}
		fmt.Println(line)
	}
	fmt.Println("end", allNewLines[len(allNewLines)-1][0])

	return allNewLines[len(allNewLines)-1][0]
}

func calculateLinesDiff(line []int) []int {
	newLine := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		newLine[i-1] = line[i] - line[i-1]
	}
	return newLine
}

// -18959
// -15413
// 403
// -59
// 913
