package utils

import (
	"flag"
	"time"
)

type Config struct {
	Year int
	Day  int

	Part1 bool
	Part2 bool

	TestInput bool
}

func ParseConfig() Config {
	now := time.Now().UTC()
	year := flag.Int("year", now.Year(), "Day")
	day := flag.Int("day", now.Day(), "Day")

	part1 := flag.Bool("part1", false, "Part 1")
	part2 := flag.Bool("part2", false, "Part 2")

	testInput := flag.Bool("test", false, "Test input")
	flag.Parse()

	if !*part1 && !*part2 {
		*part1 = true
		*part2 = true
	}

	return Config{
		Year: *year,
		Day:  *day,

		Part1: *part1,
		Part2: *part2,

		TestInput: *testInput,
	}
}
