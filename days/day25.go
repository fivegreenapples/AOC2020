package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day25Part1(in string) string {

	publicKeys := utils.LinesAsInts(in)

	loopSizes := map[int]int{}
	for idx, pk := range publicKeys[0:2] {

		loop := 0
		val := 1
		subj := 7
		for {

			loop++

			val *= subj
			val = val % 20201227

			if val == pk {
				break
			}
		}

		loopSizes[idx] = loop
	}

	val := 1
	subj := publicKeys[0]
	for loop := 0; loop < loopSizes[1]; loop++ {
		val *= subj
		val = val % 20201227
	}

	return strconv.Itoa(val)
}
