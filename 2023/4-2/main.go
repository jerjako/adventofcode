package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	chars, err := os.ReadFile("./2023/4-2/input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	lines := strings.Split(string(chars), "\n")

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

	fmt.Println("total: " + strconv.Itoa(total))
}
