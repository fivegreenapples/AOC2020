package days

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day10Part1(in string) string {

	adaptors := utils.LinesAsInts(in)
	sort.Ints(adaptors)

	num1Jolt, num3Jolt := 0, 0

	prev := 0
	for _, a := range adaptors {

		if a-prev > 3 {
			fmt.Println("found joltage difference > 3 between", prev, "and", a)
			return "?"
		}

		if a-prev == 1 {
			num1Jolt++
		} else if a-prev == 3 {
			num3Jolt++
		}

		prev = a
	}
	num3Jolt++

	return strconv.Itoa(num1Jolt * num3Jolt)
}

func (r *Runner) Day10Part2(in string) string {

	adaptors := utils.LinesAsInts(in)
	sort.Ints(adaptors)
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)

	sequencesByLength := map[int]int{}

	prev := 0
	currentSeqLength := 0
	for _, a := range adaptors {

		if a-prev > 3 {
			fmt.Println("found joltage difference > 3 between", prev, "and", a)
			return "?"
		}

		if a-prev == 1 {
			currentSeqLength++
		} else {
			if currentSeqLength > 0 {
				sequencesByLength[currentSeqLength]++
			}
			currentSeqLength = 0
		}

		prev = a
	}

	count := 1
	for len, num := range sequencesByLength {

		for ; num > 0; num-- {
			switch len {
			case 1:
			case 2:
				count = count + count*1
			case 3:
				count = count + count*3
			case 4:
				count = count + count*6
			default:
				fmt.Println("unhandled sequence length of", len)
			}
		}

	}

	return strconv.Itoa(count)
}
