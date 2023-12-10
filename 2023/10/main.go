package main

import (
	"fmt"
	"slices"
	"sync"

	"github.com/jerjako/adventofcode/utils"
)

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/

type grid map[int]map[int]string

// generateGrid will generate a 2d grid with [x][y] coordinates.
func generateGrid(lines []string) (grid, int, int) {
	g := make(grid, len(lines[0]))
	var sX, sY int
	for x, line := range lines {
		for y, tile := range line {
			if _, ok := g[y]; !ok {
				g[y] = make(map[int]string, len(lines))
			}

			tileStr := string(tile)
			g[y][x] = tileStr
			if tileStr == "S" {
				sX = y
				sY = x
			}
		}
	}

	return g, sX, sY
}

func coordToStr(x, y int) string {
	return utils.ToString(x) + "/" + utils.ToString(y)
}

func loopDistance(g grid, startX, startY, checkPointX, checkpointY int, avoid map[string]struct{}, distance int, first bool) int {
	x := checkPointX
	y := checkpointY

	for {
		if x == startX && y == startY {
			if !first {
				return distance
			}
		} else {
			// Avoid passing here again.
			avoid[coordToStr(x, y)] = struct{}{}
		}

		distance++

		possibilitiesFound := [][2]int{}

		// To North: | 7 F
		if y > 0 {
			if slices.Contains([]string{"|", "7", "F"}, g[x][y-1]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x, y - 1})
			}
		}

		// To East: - J 7
		if x < len(g[0])-1 {
			if slices.Contains([]string{"-", "J", "7"}, g[x+1][y]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x + 1, y})
			}
		}

		// To South: | L J
		if y < len(g)-1 {
			if slices.Contains([]string{"|", "L", "J"}, g[x][y+1]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x, y + 1})
			}
		}

		// To West: - L F
		if x > 0 {
			if slices.Contains([]string{"-", "L", "F"}, g[x-1][y]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x - 1, y})
			}
		}

		// Reduce the possibilities with the avoid map.
		newPossibilitiesFound := [][2]int{}
		for _, possibility := range possibilitiesFound {
			if possibility[0] == startX && possibility[1] == startY && !first {
				return distance
			}

			if _, toAvoid := avoid[coordToStr(possibility[0], possibility[1])]; !toAvoid {
				newPossibilitiesFound = append(newPossibilitiesFound, possibility)
			}
		}
		possibilitiesFound = newPossibilitiesFound

		first = false

		if len(possibilitiesFound) == 0 {
			return -1
		} else if len(possibilitiesFound) == 1 {
			x = possibilitiesFound[0][0]
			y = possibilitiesFound[0][1]
			continue
		}

		// Start a new loopDistance for each possibilities by avoiding the others.
		for _, possibility := range possibilitiesFound {
			possibleDistance := loopDistance(g, startX, startY, possibility[0], possibility[1], avoid, distance, true)
			if possibleDistance > 0 {
				return possibleDistance
			}
		}
	}
}

func part1(lines []string) string {
	distances := []int{}

	// Test with each possible S start.
	var wg sync.WaitGroup
	for _, s := range []string{"|", "-", "L", "J", "7", "F"} {
		wg.Add(1)

		go func(newS string) {
			defer wg.Done()

			g, sX, sY := generateGrid(lines)
			g[sX][sY] = newS

			d := loopDistance(g, sX, sY, sX, sY, map[string]struct{}{}, 0, true)
			distances = append(distances, d)
		}(s)
	}

	wg.Wait()

	return utils.ToString(int(slices.Max(distances) / 2))
}

func part2(lines []string) string {
	return "give up"
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
