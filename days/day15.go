package days

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day15Part1(in string) string {
	return strconv.Itoa(r.day15Process(in, 2020))
}

func (r *Runner) Day15Part2(in string) string {
	return strconv.Itoa(r.day15Process(in, 30000000))
}

func (r *Runner) day15Process(in string, maxIdx int) int {

	startingNumbers := utils.CsvToInts(in)
	n := time.Now()
	lastLastSeen := make([]int32, maxIdx)
	lastSeen := make([]int32, maxIdx)
	fmt.Println(time.Now().Sub(n))
	for idx, num := range startingNumbers {
		lastLastSeen[num] = int32(idx) + 1
		lastSeen[num] = int32(idx) + 1
	}

	prev := int32(startingNumbers[len(startingNumbers)-1])
	var thisNum int32
	for idx := len(startingNumbers); idx < maxIdx; idx++ {
		thisNum = lastSeen[prev] - lastLastSeen[prev]

		if lastSeen[thisNum] == 0 {
			lastLastSeen[thisNum] = int32(idx) + 1
			lastSeen[thisNum] = int32(idx) + 1
		} else {
			lastLastSeen[thisNum] = lastSeen[thisNum]
			lastSeen[thisNum] = int32(idx) + 1
		}
		prev = thisNum
	}

	return int(thisNum)
}
