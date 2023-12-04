package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	chars, err := os.ReadFile("./2023/4/input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	lines := strings.Split(string(chars), "\n")
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
				fmt.Println(me, winningCnt)
			}
		}
		total += winningCnt
	}

	fmt.Println("total: " + strconv.Itoa(total))
}
