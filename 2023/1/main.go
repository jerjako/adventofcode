package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	chars, err := os.ReadFile("./2023/1/input.txt")
	if err != nil {
		panic(err)
	}
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

	fmt.Println("total: " + strconv.Itoa(total))
}
