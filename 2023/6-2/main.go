package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	chars, err := os.ReadFile("./2023/6-2/input.txt")
	if err != nil {
		panic(err)
	}

	timesDistanceSrc := map[int]int{}

	r := regexp.MustCompile(`(\d+)`)
	items := r.FindAllString(strings.ReplaceAll(string(chars), " ", ""), -1)
	fmt.Println(items)
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

	fmt.Println(timesDistanceSrc)

	fmt.Println("total: ", total)
}
