package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	byteNewLine = byte(10)
	byte0       = byte(48)
	byte9       = byte(57)
	byteDot     = byte(46)
	byteGear    = byte(42) // *

	debug = false
)

func main() {
	chars, err := os.ReadFile("./2023/3-2/input.txt")
	if err != nil {
		panic(err)
	}

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

		if isNumber(c) {
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
				if coord := checkSymbolAround(m, x, y, &activeNumberValid); coord != "" {
					symbolCnt[coord]++
					activeNumberValidCoord = coord
				}
			} else if activeNumber >= 0 && n >= 0 {
				// number continue
				activeNumber = activeNumber*10 + n
				if coord := checkSymbolAround(m, x, y, &activeNumberValid); coord != "" {
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

	fmt.Println("total: " + strconv.Itoa(total))
}

func isNumber(c byte) bool {
	return c >= byte0 && c <= byte9
}

func checkSymbolAround(m map[int]map[int]int, x, y int, activeNumberValid *bool) string {
	if *activeNumberValid == true {
		return ""
	}

	if checkSymbol(m, x-1, y-1, activeNumberValid) {
		return coordConcat(x-1, y-1)
	}
	if checkSymbol(m, x-1, y, activeNumberValid) {
		return coordConcat(x-1, y)
	}
	if checkSymbol(m, x-1, y+1, activeNumberValid) {
		return coordConcat(x-1, y+1)
	}
	if checkSymbol(m, x, y-1, activeNumberValid) {
		return coordConcat(x, y-1)
	}

	if checkSymbol(m, x, y+1, activeNumberValid) {
		return coordConcat(x, y+1)
	}
	if checkSymbol(m, x+1, y-1, activeNumberValid) {
		return coordConcat(x+1, y-1)
	}
	if checkSymbol(m, x+1, y, activeNumberValid) {
		return coordConcat(x+1, y)
	}
	if checkSymbol(m, x+1, y+1, activeNumberValid) {
		return coordConcat(x+1, y+1)
	}

	return ""
}

func checkSymbol(m map[int]map[int]int, x, y int, activeNumberValid *bool) bool {
	if *activeNumberValid == true {
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

func coordConcat(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
