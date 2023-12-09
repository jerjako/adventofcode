package utils

func RunDay() ([]string, bool, bool) {
	config := ParseConfig()
	lines := ReadLines(config.Year, config.Day, config.TestInput)

	return lines, config.Part1, config.Part2
}
