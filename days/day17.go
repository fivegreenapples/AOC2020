package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day17Part1(in string) string {
	lines := utils.MultilineCsvToStrings(in, "")
	grid := map[utils.Coord3d]bool{}

	for y, line := range lines {
		for x, char := range line {
			grid[utils.Coord3d{
				X: x,
				Y: y,
				Z: 0,
			}] = char == "#"
		}
	}

	startingMin, startingMax := utils.ExtentsOf3DBoolMap(grid)
	extend := utils.Coord3d{X: 1, Y: 1, Z: 1}

	maxCycle := 6
	var c int
	for c = 1; c <= maxCycle; c++ {
		grid = r.day17CycleGrid3d(grid, startingMin.Sub(extend.Scale(c)), startingMax.Add(extend.Scale(c)))
	}

	activeCount := 0
	utils.Foreach3D(startingMin.Sub(extend.Scale(c-1)), startingMax.Add(extend.Scale(c-1)), func(pos utils.Coord3d) {
		if grid[pos] {
			activeCount++
		}
	})

	return strconv.Itoa(activeCount)
}

func (r *Runner) Day17Part2(in string) string {
	lines := utils.MultilineCsvToStrings(in, "")
	grid := map[utils.Coord4d]bool{}

	for y, line := range lines {
		for x, char := range line {
			grid[utils.Coord4d{
				X: x,
				Y: y,
				Z: 0,
				W: 0,
			}] = char == "#"
		}
	}

	startingMin, startingMax := utils.ExtentsOf4DBoolMap(grid)
	extend := utils.Coord4d{X: 1, Y: 1, Z: 1, W: 1}

	maxCycle := 6
	var c int
	for c = 1; c <= maxCycle; c++ {
		grid = r.day17CycleGrid4d(grid, startingMin.Sub(extend.Scale(c)), startingMax.Add(extend.Scale(c)))
	}

	activeCount := 0
	utils.Foreach4D(startingMin.Sub(extend.Scale(c-1)), startingMax.Add(extend.Scale(c-1)), func(pos utils.Coord4d) {
		if grid[pos] {
			activeCount++
		}
	})

	return strconv.Itoa(activeCount)
}

func (r *Runner) day17CycleGrid3d(grid map[utils.Coord3d]bool, min, max utils.Coord3d) map[utils.Coord3d]bool {
	newGrid := map[utils.Coord3d]bool{}
	utils.Foreach3D(min, max, func(pos utils.Coord3d) {

		currentActive := grid[pos]

		adjacentActiveCount := 0
		for _, adj := range pos.Adjacents() {
			if grid[adj] {
				adjacentActiveCount++
			}
		}
		newGrid[pos] = (currentActive && (adjacentActiveCount == 2 || adjacentActiveCount == 3)) ||
			(!currentActive && adjacentActiveCount == 3)
	})
	return newGrid
}
func (r *Runner) day17CycleGrid4d(grid map[utils.Coord4d]bool, min, max utils.Coord4d) map[utils.Coord4d]bool {
	newGrid := map[utils.Coord4d]bool{}
	utils.Foreach4D(min, max, func(pos utils.Coord4d) {

		currentActive := grid[pos]

		adjacentActiveCount := 0
		for _, adj := range pos.Adjacents() {
			if grid[adj] {
				adjacentActiveCount++
			}
		}
		newGrid[pos] = (currentActive && (adjacentActiveCount == 2 || adjacentActiveCount == 3)) ||
			(!currentActive && adjacentActiveCount == 3)
	})
	return newGrid
}
