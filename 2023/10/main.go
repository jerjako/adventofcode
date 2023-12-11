package main

import (
	"fmt"
	"slices"
	"strings"
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

type grid map[string]string

var (
	bestDistance = 0
	bestPath     map[string]struct{}
	bestGrid     grid
)

// generateGrid will generate a 2d grid with [x][y] coordinates.
func generateGrid(lines []string) (grid, int, int) {
	g := grid{}
	var sX, sY int
	for x, line := range lines {
		for y, tile := range line {
			tileStr := string(tile)
			g[coordToStr(x, y)] = tileStr
			if tileStr == "S" {
				sX = x
				sY = y
			}
		}
	}

	fmt.Print("\n-///--\n")
	xLen, yLen := gridGridSize(g)
	for x := 0; x <= xLen; x++ {
		for y := 0; y <= yLen; y++ {
			fmt.Print(g[coordToStr(x, y)])
		}
		fmt.Print("\n")
	}

	return g, sX, sY
}

func copyGrid(g grid) grid {
	newGrid := grid{}

	for l := range g {
		newGrid[l] = g[l]
	}

	return newGrid
}

func coordToStr(x, y int) string {
	return utils.ToString(x) + "/" + utils.ToString(y)
}

func strToCoord(str string) (int, int) {
	s := strings.Split(str, "/")
	return utils.ToInt(s[0]), utils.ToInt(s[1])
}

func copyMap(m map[string]struct{}) map[string]struct{} {
	newMap := make(map[string]struct{}, len(m))
	for i := range m {
		newMap[i] = struct{}{}
	}
	return newMap
}

func loopDistance(g grid, startX, startY int, path map[string]struct{}, first bool) map[string]struct{} {
	x := startX
	y := startY

	previousPos := [2]int{x, y}

	for {
		if x == startX && y == startY {
			if !first {
				return path
			}
		}

		coord := coordToStr(x, y)
		path[coord] = struct{}{}

		possibilitiesFound := [][2]int{}

		// To North: | 7 F
		if x > 0 && slices.Contains([]string{"|", "L", "J"}, g[coord]) {
			if slices.Contains([]string{"|", "7", "F"}, g[coordToStr(x-1, y)]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x - 1, y})
			}
		}

		// To East: - J 7
		if slices.Contains([]string{"-", "L", "F"}, g[coord]) {
			if slices.Contains([]string{"-", "J", "7"}, g[coordToStr(x, y+1)]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x, y + 1})
			}
		}

		// To South: | L J
		if slices.Contains([]string{"|", "7", "F"}, g[coord]) {
			if slices.Contains([]string{"|", "L", "J"}, g[coordToStr(x+1, y)]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x + 1, y})
			}
		}

		// To West: - L F
		if x > 0 && slices.Contains([]string{"-", "J", "7"}, g[coord]) {
			if slices.Contains([]string{"-", "L", "F"}, g[coordToStr(x, y-1)]) {
				possibilitiesFound = append(possibilitiesFound, [2]int{x, y - 1})
			}
		}

		fmt.Println(possibilitiesFound, first)
		// Reduce the possibilities with the avoid map.
		newPossibilitiesFound := [][2]int{}
		for _, possibility := range possibilitiesFound {
			if possibility[0] == startX && possibility[1] == startY && len(path) > 2 {
				return path
			}

			//if _, toAvoid := avoid[coordToStr(possibility[0], possibility[1])]; !toAvoid {
			//	newPossibilitiesFound = append(newPossibilitiesFound, possibility)
			//}

			if _, alreadyPass := path[coordToStr(possibility[0], possibility[1])]; !alreadyPass {
				newPossibilitiesFound = append(newPossibilitiesFound, possibility)
				break
			}

			//if !(previousPos[0] == possibility[0] && previousPos[1] == possibility[1]) {
			//	newPossibilitiesFound = append(newPossibilitiesFound, possibility)
			//	break
			//}
		}
		possibilitiesFound = newPossibilitiesFound

		fmt.Println(x, y, possibilitiesFound, g[coord])

		first = false

		if len(possibilitiesFound) == 0 {
			fmt.Println(" NIL")
			return nil
		} else if len(possibilitiesFound) == 1 {
			previousPos[0] = x
			previousPos[1] = y

			x = possibilitiesFound[0][0]
			y = possibilitiesFound[0][1]

			continue
		} else if first {
			continue
		}
	}
}

func calculateBestLoop(lines []string) {
	paths := []map[string]struct{}{}
	grids := []grid{}

	g, sX, sY := generateGrid(lines)

	// Test with each possible S start.
	rangeList := []string{"|", "-", "L", "J", "7", "F"}
	//rangeList = []string{"F"}
	var wg sync.WaitGroup
	for _, s := range rangeList {
		wg.Add(1)

		go func(g grid, newS string) {
			defer wg.Done()

			newGrid := copyGrid(g)
			newGrid[coordToStr(sX, sY)] = newS

			d := loopDistance(newGrid, sX, sY, map[string]struct{}{}, true)
			grids = append(grids, newGrid)
			paths = append(paths, d)
		}(g, s)
	}

	wg.Wait()

	bestDistance = 0
	for i, path := range paths {
		fmt.Println(path)
		if len(path)/2 > bestDistance {
			bestDistance = len(path) / 2
			bestPath = path
			bestGrid = grids[i]
		}
	}
}

func gridGridSize(g grid) (int, int) {
	maxXLen := 0
	maxYLen := 0
	for i := range g {
		x, y := strToCoord(i)
		if x > maxXLen {
			maxXLen = x
		}
		if y > maxYLen {
			maxYLen = y
		}
	}

	return maxXLen, maxYLen
}

func drawBestPath() {
	xLen, yLen := gridGridSize(bestGrid)
	for x := 0; x <= xLen; x++ {
		for y := 0; y <= yLen; y++ {
			fmt.Print(bestGrid[coordToStr(x, y)])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	fmt.Print("\n")
	for x := 0; x <= xLen; x++ {
		for y := 0; y <= yLen; y++ {
			if _, ok := bestPath[coordToStr(x, y)]; ok {
				fmt.Print(bestGrid[coordToStr(x, y)])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func part1(lines []string) string {
	calculateBestLoop(lines)
	return utils.ToString(bestDistance)
}

func part2(lines []string) string {
	fmt.Println(bestPath)
	fmt.Println(bestGrid)

	drawBestPath()

	//	path := copyMap(bestPath)

	total := 0
	xLen, yLen := gridGridSize(bestGrid)
	for x := 0; x <= xLen; x++ {
		in := false
		//skipBefore := false
		for y := 0; y <= yLen; y++ {
			// Vertical line
			//if in && bestGrid[x][y] == "|" {
			//	skipBefore = true
			//	continue
			//}

			if _, ok := bestPath[coordToStr(x, y)]; ok {

				//F--JF--7||LJLJ.F7FJ

				// Skip non changing pattern.
				if bestGrid[coordToStr(x, y)] == "F" || bestGrid[coordToStr(x, y)] == "L" {
					startLetter := bestGrid[coordToStr(x, y)]
					fmt.Println(startLetter, x, y)
					// Skip the dashes.
					for newY := y + 1; newY <= yLen; newY++ {
						if _, okPath := bestPath[coordToStr(x, newY)]; okPath {
							fmt.Println(bestGrid[coordToStr(x, newY)], x, newY)
							if bestGrid[coordToStr(x, newY)] == "-" {
								fmt.Println("skipped", x, newY)
								y = newY
								continue
							} else {
								break
							}
						} else {
							break
						}
					}
					fmt.Println(startLetter, bestGrid[coordToStr(x, y+1)])
					if (startLetter == "F" && bestGrid[coordToStr(x, y+1)] == "7") ||
						(startLetter == "L" && bestGrid[coordToStr(x, y+1)] == "J") {
						y++
						continue
					} else {
						y++
					}
				}

				/*if in {
					if _, willBeAfter := bestPath[coordToStr(x, y+1)]; willBeAfter {
						in = !in
						continue
					}

					if _, wasBefore := bestPath[coordToStr(x, y-1)]; wasBefore {
						continue
					}
				}*/

				in = !in
				fmt.Println("CHANGE IN", x, y, in)
				//fmt.Println(x, y, "change IN", in)

				//	isSkipped := false
				// For horizontal lines, skip them.
				/*
					if slices.Contains([]string{"L", "F"}, bestGrid[coordToStr(x, y)]) {
						for newY := y + 1; newY <= yLen; newY++ {
							//	fmt.Println(bestGrid[coordToStr(x, newY)])
							if _, okPath := bestPath[coordToStr(x, newY)]; okPath {
								isSkipped = true
								if !slices.Contains([]string{"-"}, bestGrid[coordToStr(x, newY)]) {
									y = newY
									break
								}
							} else {
								panic("impossible")
								y = newY
								break
							}
						}
					}
				*/

				/*for newY := y + 1; newY <= yLen; newY++ {
					if _, okPath := bestPath[coordToStr(x, newY)]; okPath {
						isSkipped = true
						y = newY
					} else {
						break
					}
				}

				// We skiped to the end of the grid.
				if y > yLen {
					break
				}
				if isSkipped {
					continue
				}*/
			} else if in {
				fmt.Println("-----COUNT FOR ", x, y)
				total++
			}

			// check end of line
			/*if in {
				isEOL := true
				for newY := y + 1; newY <= yLen; newY++ {
					if _, okPath := bestPath[coordToStr(x, newY)]; okPath {
						isEOL = false
						break
					}
				}
				if isEOL {
					y = yLen
					in = false
					continue
				}
			}*/

			//skipBefore = false
			/*_, inPath := bestPath[coordToStr(x, y)]
			fmt.Println(x, y, bestGrid[coordToStr(x, y)], inPath, in, (!inPath && in), in && !slices.Contains([]string{"|", "-", "L", "J", "7", "F"}, bestGrid[coordToStr(x, y)]))
			if !inPath && in {
				total++
			}*/
		}
	}

	return utils.ToString(total)
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

// p1
// 6831

// p2
// 2851 too high
// 2847 too high
// 2721 too high
// 917
// 160
// 672
