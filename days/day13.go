package days

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day13Part1(in string) string {

	inputs := utils.Lines(in)

	earliestTS := utils.MustAtoi(inputs[0])
	buses := utils.CsvToStrings(inputs[1])

	earliestBus, waitMinutes := 0, math.MaxInt64
	for _, bus := range buses {
		if bus == "x" {
			continue
		}

		busId := utils.MustAtoi(bus)

		missMinutes := earliestTS % busId
		if missMinutes == 0 {
			return "0"
		}
		thisWait := busId - missMinutes
		if thisWait < waitMinutes {
			waitMinutes = thisWait
			earliestBus = busId
		}

	}

	return strconv.Itoa(earliestBus * waitMinutes)
}

func (r *Runner) Day13Part2(in string) string {

	inputs := utils.Lines(in)

	buses := utils.CsvToStrings(inputs[1])

	if r.verbose {
		fmt.Println(buses)
	}
	combinedInterval := utils.MustAtoi(buses[0])
	combinedOffset := 0
	for thisOffset, bus := range buses[1:] {

		thisOffset++
		if bus == "x" {
			continue
		}
		busId := utils.MustAtoi(bus)

		cycles := 1
		for (combinedOffset+(cycles*combinedInterval)+thisOffset)%busId != 0 {
			cycles++
		}

		combinedOffset = combinedOffset + cycles*combinedInterval
		combinedInterval = utils.LowestCommonMultiple(combinedInterval, busId)
		if r.verbose {
			fmt.Println(busId, cycles, combinedOffset, combinedInterval)
		}

	}
	return strconv.Itoa(combinedOffset)
}
