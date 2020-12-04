package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day3Part1(in string) string {
	mapLines := utils.MultilineCsvToStrings(in, "")
	numTrees := r.day3CalcTrees(mapLines, utils.Coord{X: 3, Y: 1})
	return strconv.Itoa(numTrees)
}

func (r *Runner) Day3Part2(in string) string {
	mapLines := utils.MultilineCsvToStrings(in, "")
	numTrees := r.day3CalcTrees(mapLines, utils.Coord{X: 1, Y: 1})
	numTrees *= r.day3CalcTrees(mapLines, utils.Coord{X: 3, Y: 1})
	numTrees *= r.day3CalcTrees(mapLines, utils.Coord{X: 5, Y: 1})
	numTrees *= r.day3CalcTrees(mapLines, utils.Coord{X: 7, Y: 1})
	numTrees *= r.day3CalcTrees(mapLines, utils.Coord{X: 1, Y: 2})
	return strconv.Itoa(numTrees)
}

func (r *Runner) day3CalcTrees(mapLines [][]string, slope utils.Coord) int {
	move := slope
	pos := utils.Coord{X: 0, Y: 0}
	numTrees := 0

	for pos.Y < len(mapLines) {

		thisX := pos.X % len(mapLines[pos.Y])
		if mapLines[pos.Y][thisX] == "#" {
			numTrees++
		}

		pos = pos.Add(move)

	}
	return numTrees
}
