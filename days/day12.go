package days

import (
	"fmt"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day12Part1(in string) string {

	instructions := utils.StringsFromRegex(in, `^([NSEWLRF])([0-9]+)$`)

	pos := utils.Coord{}
	dir := utils.Right
	for _, instr := range instructions {

		action := instr[1]
		value := utils.MustAtoi(instr[2])

		switch action {
		case "N":
			pos = pos.Add(utils.Up.Scale(value))
		case "S":
			pos = pos.Add(utils.Down.Scale(value))
		case "E":
			pos = pos.Add(utils.Right.Scale(value))
		case "W":
			pos = pos.Add(utils.Left.Scale(value))
		case "L":
			dir = dir.RotateAntiClockwise(value)
		case "R":
			dir = dir.RotateClockwise(value)
		case "F":
			pos = pos.Add(dir.Scale(value))
		}
		if r.verbose {
			fmt.Println("pos:", pos, "dir:", dir)
		}
	}
	return strconv.Itoa(pos.Manhattan())
}

func (r *Runner) Day12Part2(in string) string {

	instructions := utils.StringsFromRegex(in, `^([NSEWLRF])([0-9]+)$`)

	shipPos := utils.Coord{}
	wpOffset := utils.Coord{X: 10, Y: -1}
	for _, instr := range instructions {

		action := instr[1]
		value := utils.MustAtoi(instr[2])

		switch action {
		case "N":
			wpOffset = wpOffset.Add(utils.Up.Scale(value))
		case "S":
			wpOffset = wpOffset.Add(utils.Down.Scale(value))
		case "E":
			wpOffset = wpOffset.Add(utils.Right.Scale(value))
		case "W":
			wpOffset = wpOffset.Add(utils.Left.Scale(value))
		case "L":
			wpOffset = wpOffset.RotateAntiClockwise(value)
		case "R":
			wpOffset = wpOffset.RotateClockwise(value)
		case "F":
			shipPos = shipPos.Add(wpOffset.Scale(value))
		}
		if r.verbose {
			fmt.Println("shipPos:", shipPos, "wpOffset:", wpOffset)
		}
	}
	return strconv.Itoa(shipPos.Manhattan())
}
