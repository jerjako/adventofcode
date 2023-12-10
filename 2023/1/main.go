package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jerjako/adventofcode/utils"
)

func part1(lines []string) string {
	// Convert from bytes to strings to bytes because of late refactor and too lazy to redo it clean.
	chars := []byte(strings.Join(lines, "\n"))

	// append a new line
	chars = append(chars, byte(10))

	total := 0
	calibrationStart := -1
	calibrationEnd := -1
	for _, char := range chars {
		if char == 10 {
			if calibrationStart == -1 {
				panic("no number found in line")
			}
			if calibrationEnd == -1 {
				calibrationEnd = calibrationStart
			}
			total += (calibrationStart*10 + calibrationEnd)

			calibrationStart = -1
			calibrationEnd = -1
		}

		if char >= 48 && char <= 57 {
			if calibrationStart == -1 {
				calibrationStart = int(char) - 48
			} else {
				calibrationEnd = int(char) - 48
			}
		}
	}

	return ("total: " + utils.ToString(total))
}

const (
	newLine = byte(10)
	byte0   = byte(48)
	byte9   = byte(57)
)

var (
	byteTextNumber = [][]byte{
		[]byte("one"),
		[]byte("two"),
		[]byte("three"),
		[]byte("four"),
		[]byte("five"),
		[]byte("six"),
		[]byte("seven"),
		[]byte("eight"),
		[]byte("nine"),
	}
)

func part2(lines []string) string {
	// Convert from bytes to strings to bytes because of late refactor and too lazy to redo it clean.
	chars := []byte(strings.Join(lines, "\n"))

	// Append a new line at the end of the inputs if needed.
	if chars[len(chars)-1] != newLine {
		chars = append(chars, newLine)
	}

	intByte0 := int(byte0)
	total := 0
	calibrationStart := -1
	calibrationEnd := -1
	last6bytes := []byte{}
	for _, char := range chars {
		// Keep the last 6 bytes.
		if len(last6bytes) == 6 {
			last6bytes = last6bytes[1:]
		}
		last6bytes = append(last6bytes, char)

		if char == newLine {
			if calibrationStart == -1 {
				panic("no number found in line")
			}
			if calibrationEnd == -1 {
				calibrationEnd = calibrationStart
			}
			// fmt.Println("line end: ", calibrationStart, calibrationEnd)
			total += (calibrationStart*10 + calibrationEnd)

			calibrationStart = -1
			calibrationEnd = -1
		}

		// Check if the last 6 bytes are not a number.
		for textNumberIndex, textNumber := range byteTextNumber {
			if len(last6bytes) < len(textNumber) {
				continue
			}
			if slices.Compare(textNumber, last6bytes[len(last6bytes)-len(textNumber):]) == 0 {
				if calibrationStart == -1 {
					calibrationStart = textNumberIndex + 1
				} else {
					calibrationEnd = textNumberIndex + 1
				}
			}
		}

		// Check if the char is a number.
		if char >= byte0 && char <= byte9 {
			if calibrationStart == -1 {
				calibrationStart = int(char) - intByte0
			} else {
				calibrationEnd = int(char) - intByte0
			}
		}
	}

	return "total: " + utils.ToString(total)
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
