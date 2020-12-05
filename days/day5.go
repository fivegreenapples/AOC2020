package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day5Part1(in string) string {

	passes := utils.Lines(in)

	maxSeat := math.MinInt64
	for _, pp := range passes {

		p := pp
		p = strings.ReplaceAll(p, "B", "1")
		p = strings.ReplaceAll(p, "F", "0")
		p = strings.ReplaceAll(p, "R", "1")
		p = strings.ReplaceAll(p, "L", "0")

		seatID, err := strconv.ParseInt(p, 2, 64)
		if err != nil {
			panic(fmt.Sprintf("failed calculating seat id for boarding pass: %s, converted val: %s, error: %s", pp, p, err.Error()))
		}

		if int(seatID) > maxSeat {
			maxSeat = int(seatID)
		}
	}

	return strconv.Itoa(maxSeat)
}

func (r *Runner) Day5Part2(in string) string {

	passes := utils.Lines(in)

	maxSeat := math.MinInt64
	minSeat := math.MaxInt64
	allSeats := map[int]bool{}
	for _, pp := range passes {

		p := pp
		p = strings.ReplaceAll(p, "B", "1")
		p = strings.ReplaceAll(p, "F", "0")
		p = strings.ReplaceAll(p, "R", "1")
		p = strings.ReplaceAll(p, "L", "0")

		seatID, err := strconv.ParseInt(p, 2, 64)
		if err != nil {
			panic(fmt.Sprintf("failed calculating seat id for boarding pass: %s, converted val: %s, error: %s", pp, p, err.Error()))
		}

		allSeats[int(seatID)] = true
		if int(seatID) > maxSeat {
			maxSeat = int(seatID)
		}
		if int(seatID) < minSeat {
			minSeat = int(seatID)
		}
	}

	missingSeat := -1
	for s := minSeat; s <= maxSeat; s++ {
		if !allSeats[s] {
			missingSeat = s
			break
		}
	}

	return strconv.Itoa(missingSeat)
}
