package days

import (
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day6Part1(in string) string {

	groups := strings.Split(in, "\n\n")
	totalCount := 0

	for _, g := range groups {

		groupAnswered := map[rune]bool{}
		answers := utils.Lines(g)
		for _, a := range answers {
			for _, rne := range a {
				groupAnswered[rne] = true
			}
		}

		totalCount += len(groupAnswered)
	}

	return strconv.Itoa(totalCount)

}

func (r *Runner) Day6Part2(in string) string {

	groups := strings.Split(in, "\n\n")
	totalCount := 0

	for _, g := range groups {

		groupAnswered := map[rune]int{}
		answers := utils.Lines(g)
		for _, a := range answers {
			for _, rne := range a {
				groupAnswered[rne] += 1
			}
		}
		for _, val := range groupAnswered {
			if val == len(answers) {
				totalCount++
			}
		}

	}

	return strconv.Itoa(totalCount)

}
