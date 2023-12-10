package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

const (
	byteNewLine = byte(10)
	byte0       = byte(48)
	byte9       = byte(57)
	byteDot     = byte(46)
	byteGear    = byte(42) // *

	debug = false
)

func part1(lines []string) string {
	// Convert from bytes to strings to bytes because of late refactor and too lazy to redo it clean.
	chars := []byte(strings.Join(lines, "\n"))

	total := 0
	x, y := 0, 0
	m := map[int]map[int]int{}
	for _, c := range chars {
		if c == byteNewLine {
			y = 0
			x++
			continue
		}

		if _, ok := m[x]; !ok {
			m[x] = map[int]int{}
		}

		if isNumberPart1(c) {
			m[x][y] = int(c - byte0)
		} else if c == byteDot {
			m[x][y] = -1
		} else {
			m[x][y] = -2
		}

		y++
	}

	activeNumber := -1
	activeNumberValid := false
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			if debug {
				fmt.Print(strconv.Itoa(m[x][y]) + " ")
			}

			n := m[x][y]

			if activeNumber < 0 && n >= 0 {
				// first number seen
				activeNumber = n
				checkSymbolAroundPart1(m, x, y, &activeNumberValid)
			} else if activeNumber >= 0 && n >= 0 {
				// number continue
				activeNumber = activeNumber*10 + n
				checkSymbolAroundPart1(m, x, y, &activeNumberValid)
			} else if activeNumber >= 0 && n < 0 {
				// active number end, is it a valid part?
				if activeNumberValid {
					total += activeNumber
				}

				activeNumber = -1
				activeNumberValid = false
			}
		}
		if debug {
			fmt.Print("\n")
		}
	}

	return "total: " + utils.ToString(total)
}

func part2(lines []string) string {
	// Convert from bytes to strings to bytes because of late refactor and too lazy to redo it clean.
	chars := []byte(strings.Join(lines, "\n"))

	total := 0
	x, y := 0, 0
	m := map[int]map[int]int{}
	for _, c := range chars {
		if c == byteNewLine {
			y = 0
			x++
			continue
		}

		if _, ok := m[x]; !ok {
			m[x] = map[int]int{}
		}

		if isNumberPart2(c) {
			m[x][y] = int(c - byte0)
		} else if c == byteDot {
			m[x][y] = -1
		} else if c == byteGear {
			m[x][y] = -2
		}

		y++
	}

	activeNumber := -1
	activeNumberValid := false
	activeNumberValidCoord := ""
	symbolCnt := map[string]int{}
	activeCoordFirst := map[string]int{}
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			if debug {
				fmt.Print(strconv.Itoa(m[x][y]) + " ")
			}

			n := m[x][y]

			if activeNumber < 0 && n >= 0 {
				// first number seen
				activeNumber = n
				if coord := checkSymbolAroundPart2(m, x, y, &activeNumberValid); coord != "" {
					symbolCnt[coord]++
					activeNumberValidCoord = coord
				}
			} else if activeNumber >= 0 && n >= 0 {
				// number continue
				activeNumber = activeNumber*10 + n
				if coord := checkSymbolAroundPart2(m, x, y, &activeNumberValid); coord != "" {
					symbolCnt[coord]++
					activeNumberValidCoord = coord
				}
			} else if activeNumber >= 0 && n < 0 {
				// active number end, is it a valid part?

				if activeNumberValid && symbolCnt[activeNumberValidCoord] == 1 {
					activeCoordFirst[activeNumberValidCoord] = activeNumber
				} else if activeNumberValid && symbolCnt[activeNumberValidCoord] == 2 {
					total += activeNumber * activeCoordFirst[activeNumberValidCoord]
				}

				activeNumber = -1
				activeNumberValid = false
			}
		}
		if debug {
			fmt.Print("\n")
		}
	}

	return "total: " + utils.ToString(total)
}

func isNumberPart1(c byte) bool {
	return c >= byte0 && c <= byte9
}

func checkSymbolAroundPart1(m map[int]map[int]int, x, y int, activeNumberValid *bool) {
	if *activeNumberValid {
		return
	}

	checkSymbolPart1(m, x-1, y-1, activeNumberValid)
	checkSymbolPart1(m, x-1, y, activeNumberValid)
	checkSymbolPart1(m, x-1, y+1, activeNumberValid)
	checkSymbolPart1(m, x, y-1, activeNumberValid)

	checkSymbolPart1(m, x, y+1, activeNumberValid)
	checkSymbolPart1(m, x+1, y-1, activeNumberValid)
	checkSymbolPart1(m, x+1, y, activeNumberValid)
	checkSymbolPart1(m, x+1, y+1, activeNumberValid)
}

func checkSymbolPart1(m map[int]map[int]int, x, y int, activeNumberValid *bool) {
	if *activeNumberValid {
		return
	}

	if v, ok := m[x][y]; ok {
		if v == -2 {
			*activeNumberValid = true
			return
		}
	}
}

func isNumberPart2(c byte) bool {
	return c >= byte0 && c <= byte9
}

func checkSymbolAroundPart2(m map[int]map[int]int, x, y int, activeNumberValid *bool) string {
	if *activeNumberValid {
		return ""
	}

	if checkSymbolPart2(m, x-1, y-1, activeNumberValid) {
		return coordConcatPart2(x-1, y-1)
	}
	if checkSymbolPart2(m, x-1, y, activeNumberValid) {
		return coordConcatPart2(x-1, y)
	}
	if checkSymbolPart2(m, x-1, y+1, activeNumberValid) {
		return coordConcatPart2(x-1, y+1)
	}
	if checkSymbolPart2(m, x, y-1, activeNumberValid) {
		return coordConcatPart2(x, y-1)
	}

	if checkSymbolPart2(m, x, y+1, activeNumberValid) {
		return coordConcatPart2(x, y+1)
	}
	if checkSymbolPart2(m, x+1, y-1, activeNumberValid) {
		return coordConcatPart2(x+1, y-1)
	}
	if checkSymbolPart2(m, x+1, y, activeNumberValid) {
		return coordConcatPart2(x+1, y)
	}
	if checkSymbolPart2(m, x+1, y+1, activeNumberValid) {
		return coordConcatPart2(x+1, y+1)
	}

	return ""
}

func checkSymbolPart2(m map[int]map[int]int, x, y int, activeNumberValid *bool) bool {
	if *activeNumberValid {
		return true
	}

	if v, ok := m[x][y]; ok {
		if v == -2 {
			*activeNumberValid = true
			return true
		}
	}

	return false
}

func coordConcatPart2(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
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
