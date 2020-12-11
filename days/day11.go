package days

import (
	"fmt"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

const (
	day11Floor    = "."
	day11Seat     = "L"
	day11Occupied = "#"
)

func (r *Runner) Day11Part1(in string) string {

	floorplan := r.day11GetFloorplan(in)
	minFloor, maxFloor := utils.ExtentsOfStringMap(floorplan)

	for {
		numOccupied := 0
		anyChanges := false
		newFloorplan := map[utils.Coord]string{}

		if r.verbose {
			r.day11PrintFloor(floorplan)
		}
		for x := minFloor.X; x <= maxFloor.X; x++ {
			for y := minFloor.Y; y <= maxFloor.Y; y++ {

				pos := utils.Coord{X: x, Y: y}
				current := floorplan[pos]
				newFloorplan[pos] = current

				if current == day11Floor {
					continue
				}

				numAdjacentOccupied := 0
				for _, adjPos := range pos.Adjacents() {
					if floorplan[adjPos] == day11Occupied {
						numAdjacentOccupied++
					}
				}

				if current == day11Occupied {
					numOccupied++
					if numAdjacentOccupied >= 4 {
						newFloorplan[pos] = day11Seat
						anyChanges = true
					}
				} else if current == day11Seat && numAdjacentOccupied == 0 {
					newFloorplan[pos] = day11Occupied
					anyChanges = true
				}
			}
		}
		if !anyChanges {
			return strconv.Itoa(numOccupied)
		}
		floorplan = newFloorplan
	}

}

func (r *Runner) Day11Part2(in string) string {

	floorplan := r.day11GetFloorplan(in)
	minFloor, maxFloor := utils.ExtentsOfStringMap(floorplan)

	for {
		numOccupied := 0
		anyChanges := false
		newFloorplan := map[utils.Coord]string{}

		if r.verbose {
			r.day11PrintFloor(floorplan)
		}
		for x := minFloor.X; x <= maxFloor.X; x++ {
			for y := minFloor.Y; y <= maxFloor.Y; y++ {

				pos := utils.Coord{X: x, Y: y}
				current := floorplan[pos]
				newFloorplan[pos] = current

				if current == day11Floor {
					continue
				}

				numVisibleOccupied := 0
				for _, dir := range utils.CardinalsAndOrdinals {
					thisPos := pos.Add(dir)
					for thisPos.IsInside(minFloor, maxFloor) {
						if floorplan[thisPos] != day11Floor {
							if floorplan[thisPos] == day11Occupied {
								numVisibleOccupied++
							}
							break
						}
						thisPos = thisPos.Add(dir)
					}
				}

				if current == day11Occupied {
					numOccupied++
					if numVisibleOccupied >= 5 {
						newFloorplan[pos] = day11Seat
						anyChanges = true
					}
				} else if current == day11Seat && numVisibleOccupied == 0 {
					newFloorplan[pos] = day11Occupied
					anyChanges = true
				}
			}
		}
		if !anyChanges {
			return strconv.Itoa(numOccupied)
		}
		floorplan = newFloorplan
	}
}

func (r *Runner) day11GetFloorplan(in string) map[utils.Coord]string {
	plan := map[utils.Coord]string{}
	rows := utils.MultilineCsvToStrings(in, "")
	for y, row := range rows {
		for x, typ := range row {
			plan[utils.Coord{X: x, Y: y}] = typ
		}
	}

	return plan
}

func (r *Runner) day11PrintFloor(floorplan map[utils.Coord]string) {
	minFloor, maxFloor := utils.ExtentsOfStringMap(floorplan)
	for y := minFloor.Y; y <= maxFloor.Y; y++ {
		for x := minFloor.X; x <= maxFloor.X; x++ {
			fmt.Print(floorplan[utils.Coord{X: x, Y: y}])
		}
		fmt.Println()
	}
	fmt.Println()
}
