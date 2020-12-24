package days

import (
	"fmt"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day24Part1(in string) string {

	grid := r.d24ProcessInputAndSetupTiles(in)

	numBlack := 0
	for _, t := range grid {
		if t {
			numBlack++
		}
	}

	return strconv.Itoa(numBlack)
}

func (r *Runner) Day24Part2(in string) string {

	grid := r.d24ProcessInputAndSetupTiles(in)

	for i := 0; i < 100; i++ {

		newGrid := map[utils.Coord]bool{}
		// first loop to calculate black tiles and find all whites in scope
		whiteTiles := map[utils.Coord]bool{}
		for pos, t := range grid {
			if t {
				adjacents := r.d24HexagonallyAdjacent(pos)
				numAdjacentBlack := 0
				for _, adj := range adjacents {
					if grid[adj] {
						numAdjacentBlack++
					} else {
						whiteTiles[adj] = true
					}
				}
				if numAdjacentBlack == 1 || numAdjacentBlack == 2 {
					newGrid[pos] = true
				}
			}
		}
		// second loop calculates the changes to the white tiles
		for pos := range whiteTiles {
			adjacents := r.d24HexagonallyAdjacent(pos)
			numAdjacentBlack := 0
			for _, adj := range adjacents {
				if grid[adj] {
					numAdjacentBlack++
				}
			}
			if numAdjacentBlack == 2 {
				newGrid[pos] = true
			}
		}
		grid = newGrid

	}

	numBlack := 0
	for _, t := range grid {
		if t {
			numBlack++
		}
	}

	return strconv.Itoa(numBlack)
}

var (
	d24E  = utils.Coord{X: 2, Y: 0}
	d24W  = utils.Coord{X: -2, Y: 0}
	d24NE = utils.Coord{X: 1, Y: 2}
	d24SE = utils.Coord{X: 1, Y: -2}
	d24SW = utils.Coord{X: -1, Y: -2}
	d24NW = utils.Coord{X: -1, Y: 2}
)

func (r *Runner) d24ProcessInputAndSetupTiles(in string) map[utils.Coord]bool {
	tiles := utils.Lines(in)

	grid := map[utils.Coord]bool{}

	for _, t := range tiles {
		pos := utils.Coord{}
		for i := 0; i < len(t); i++ {
			switch t[i] {
			case 'e':
				pos = pos.Add(utils.Coord{X: 2, Y: 0})
			case 'w':
				pos = pos.Add(utils.Coord{X: -2, Y: 0})
			case 'n', 's':
				dir := string(t[i]) + string(t[i+1])
				i++
				switch dir {
				case "ne":
					pos = pos.Add(utils.Coord{X: 1, Y: 2})
				case "se":
					pos = pos.Add(utils.Coord{X: 1, Y: -2})
				case "sw":
					pos = pos.Add(utils.Coord{X: -1, Y: -2})
				case "nw":
					pos = pos.Add(utils.Coord{X: -1, Y: 2})
				}
			}
		}
		if r.verbose {
			fmt.Println(pos, t)
		}
		grid[pos] = !grid[pos]
	}

	return grid

}

func (r *Runner) d24HexagonallyAdjacent(pos utils.Coord) []utils.Coord {
	return []utils.Coord{
		pos.Add(utils.Coord{X: 2, Y: 0}),
		pos.Add(utils.Coord{X: -2, Y: 0}),
		pos.Add(utils.Coord{X: 1, Y: 2}),
		pos.Add(utils.Coord{X: 1, Y: -2}),
		pos.Add(utils.Coord{X: -1, Y: -2}),
		pos.Add(utils.Coord{X: -1, Y: 2}),
	}
}
