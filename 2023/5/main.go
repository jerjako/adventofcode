package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jerjako/adventofcode/utils"
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

func part1(lines []string) string {
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

	return "closest: " + utils.ToString(int(slices.Min(data)))
}

func part2(lines []string) string {
	// Add a new line at the end to trigger a last loop for append
	lines = append(lines, "")

	data := map[int]int{}
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

			for i := 0; i < len(l); i += 2 {
				seedIntStart, _ := strconv.Atoi(l[i])
				seedIntRange, _ := strconv.Atoi(l[i+1])

				data[seedIntStart] = seedIntRange
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

	lowest := int(0)

	for dataStart, dataRange := range data {
		subData := make([]int, dataRange)

		for i := int(0); i < dataRange; i++ {
			subData[i] = dataStart + i
		}
		for _, convertMap := range convertsMap {
			for i, d := range subData {
				for _, convertMapRange := range convertMap.ranges {
					if d >= convertMapRange.sourceRangeStart && d < convertMapRange.sourceRangeStart+convertMapRange.rangeLen {
						subData[i] = d - convertMapRange.sourceRangeStart + convertMapRange.destinationRangeStart
						break
					}
				}
			}
		}

		newLowest := slices.Min(subData)
		if lowest == 0 || newLowest < lowest {
			lowest = newLowest
		}
	}

	return "closest: " + utils.ToString(int(lowest))
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
