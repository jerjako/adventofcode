package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type convertMapRange struct {
	destinationRangeStart uint64
	sourceRangeStart      uint64
	rangeLen              uint64
}

type convertMap struct {
	fromMap string
	toMap   string

	ranges []convertMapRange
}

func main() {
	chars, err := os.ReadFile("./2023/5-2/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(chars), "\n")
	// Add a new line at the end to trigger a last loop for append
	lines = append(lines, "")

	data := map[uint64]uint64{}
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
				seedIntStart, _ := strconv.ParseUint(l[i], 10, 64)
				seedIntRange, _ := strconv.ParseUint(l[i+1], 10, 64)

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
			destinationRangeStart, _ := strconv.ParseUint(l[0], 10, 64)
			sourceRangeStart, _ := strconv.ParseUint(l[1], 10, 64)
			rangeLen, _ := strconv.ParseUint(l[2], 10, 64)

			convertMapCurrent.ranges = append(convertMapCurrent.ranges, convertMapRange{
				destinationRangeStart: destinationRangeStart,
				sourceRangeStart:      sourceRangeStart,
				rangeLen:              rangeLen,
			})
		}
	}

	lowest := uint64(0)

	for dataStart, dataRange := range data {
		subData := make([]uint64, dataRange)

		for i := uint64(0); i < dataRange; i++ {
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
		subData = []uint64{}
	}

	fmt.Println("closest: ", lowest)
}
