package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {
	timesDistanceSrc := map[int]int{}

	r := regexp.MustCompile(`(\d+)`)
	items := r.FindAllString(strings.Join(lines, "\n"), -1)
	for i := 0; i < len(items)/2; i++ {
		t, _ := strconv.Atoi(items[i])
		d, _ := strconv.Atoi(items[len(items)/2+i])
		timesDistanceSrc[t] = d
	}

	total := 1
	for timeTotal, distanceGoal := range timesDistanceSrc {
		subTotal := 0
		// pushDistanceTmp := make(map[int]int, timeTotal)

		for push := 0; push <= timeTotal; push++ {
			timeRemain := timeTotal - push
			distance := timeRemain * push
			// pushDistanceTmp[push] = distance

			if distance > distanceGoal {
				subTotal++
			}
		}
		total *= subTotal

	}

	return utils.ToString(total)
}

func part2(lines []string) string {

	timesDistanceSrc := map[int]int{}

	r := regexp.MustCompile(`(\d+)`)
	items := r.FindAllString(strings.ReplaceAll(strings.Join(lines, "\n"), " ", ""), -1)
	for i := 0; i < len(items)/2; i++ {
		t, _ := strconv.Atoi(items[i])
		d, _ := strconv.Atoi(items[len(items)/2+i])
		timesDistanceSrc[t] = d
	}

	total := 1
	for timeTotal, distanceGoal := range timesDistanceSrc {
		subTotal := 0
		// pushDistanceTmp := make(map[int]int, timeTotal)

		for push := 0; push <= timeTotal; push++ {
			timeRemain := timeTotal - push
			distance := timeRemain * push
			// pushDistanceTmp[push] = distance

			if distance > distanceGoal {
				subTotal++
			}
		}
		total *= subTotal

	}

	return utils.ToString(total)
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
