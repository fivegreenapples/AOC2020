package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day1Part1(in string) string {
	entries := utils.LinesAsInts(in)
	for i, valI := range entries {
		for _, valJ := range entries[i+1:] {
			if valI+valJ == 2020 {
				return strconv.Itoa(valI * valJ)
			}
		}
	}
	return ""
}

func (r *Runner) Day1Part2(in string) string {
	entries := utils.LinesAsInts(in)
	for i, valI := range entries {
		for j, valJ := range entries[i+1:] {
			for _, valK := range entries[j+1:] {
				if valI+valJ+valK == 2020 {
					return strconv.Itoa(valI * valJ * valK)
				}
			}
		}
	}
	return ""
}
