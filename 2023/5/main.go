package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type convertMapRange struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLen              int
}

type convertMap struct {
	fromMap string
	toMap   string

	ranges []convertMapRange
}

func main() {
	chars, err := os.ReadFile("./2023/5/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(chars), "\n")
	// Add a new line at the end to trigger a last loop for append
	lines = append(lines, "")

	data := []int{}
	convertsMap := []convertMap{}
	var convertMapCurrent convertMap
	for _, line := range lines {
		if line == "" {
			if convertMapCurrent.fromMap != "" {
				convertsMap = append(convertsMap, convertMapCurrent)
			}
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			l := strings.Split(strings.Split(line, ": ")[1], " ")
			for _, seed := range l {
				seedInt, _ := strconv.Atoi(seed)
				data = append(data, seedInt)
			}
			continue
		}

		if strings.HasSuffix(line, ":") {
			l := strings.Split(strings.Split(line, " ")[0], "-")
			convertMapCurrent = convertMap{
				fromMap: l[0],
				toMap:   l[2],
				ranges:  []convertMapRange{},
			}
			continue
		} else {
			l := strings.Split(line, " ")
			destinationRangeStart, _ := strconv.Atoi(l[0])
			sourceRangeStart, _ := strconv.Atoi(l[1])
			rangeLen, _ := strconv.Atoi(l[2])

			convertMapCurrent.ranges = append(convertMapCurrent.ranges, convertMapRange{
				destinationRangeStart: destinationRangeStart,
				sourceRangeStart:      sourceRangeStart,
				rangeLen:              rangeLen,
			})
		}
	}

	for _, convertMap := range convertsMap {
		newData := make([]int, len(data))
		for i, d := range data {
			rangeFound := false
			for _, convertMapRange := range convertMap.ranges {
				if d >= convertMapRange.sourceRangeStart && d <= convertMapRange.sourceRangeStart+convertMapRange.rangeLen {
					newData[i] = d - convertMapRange.sourceRangeStart + convertMapRange.destinationRangeStart
					rangeFound = true
					break
				}
			}

			if !rangeFound {
				newData[i] = d
			}
		}
		data = newData
	}

	fmt.Println("closest: ", slices.Min(data))
}
