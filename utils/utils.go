package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(year, day int, testInput bool) []string {
	filename := "./" + strconv.Itoa(year) + "/" + strconv.Itoa(day) + "/"
	if testInput {
		filename += "test.txt"
	} else {
		filename += "input.txt"
	}
	chars, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(chars), "\n")
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func ToString(i int) string {
	return strconv.Itoa(i)
}
