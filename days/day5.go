package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day5Part1(in string) string {

	in = strings.ReplaceAll(in, "B", "1")
	in = strings.ReplaceAll(in, "F", "0")
	in = strings.ReplaceAll(in, "R", "1")
	in = strings.ReplaceAll(in, "L", "0")
	passes := utils.Lines(in)

	maxSeat := math.MinInt64
	for _, p := range passes {
		seatID, err := strconv.ParseInt(p, 2, 64)
		if err != nil {
			panic(fmt.Sprintf("failed calculating seat id for boarding pass %s: %s", p, err.Error()))
		}

		if int(seatID) > maxSeat {
			maxSeat = int(seatID)
		}
	}

	return strconv.Itoa(maxSeat)
}

func (r *Runner) Day5Part2(in string) string {

	in = strings.ReplaceAll(in, "B", "1")
	in = strings.ReplaceAll(in, "F", "0")
	in = strings.ReplaceAll(in, "R", "1")
	in = strings.ReplaceAll(in, "L", "0")
	passes := utils.Lines(in)

	maxSeat := math.MinInt64
	minSeat := math.MaxInt64
	allSeats := map[int]bool{}
	for _, p := range passes {
		seatID, err := strconv.ParseInt(p, 2, 64)
		if err != nil {
			panic(fmt.Sprintf("failed calculating seat id for boarding pass %s: %s", p, err.Error()))
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
