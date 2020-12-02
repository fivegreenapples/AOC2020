package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day2Part1(in string) string {
	passwordData := utils.StringsFromRegex(in, `^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$`)

	valid := 0
	for _, p := range passwordData {
		min := utils.MustAtoi(p[1])
		max := utils.MustAtoi(p[2])
		char := rune(p[3][0])
		password := p[4]

		counts := day2StringToCharCounts(password)
		if counts[char] >= min && counts[char] <= max {
			valid++
		}
	}
	return strconv.Itoa(valid)
}

func (r *Runner) Day2Part2(in string) string {
	passwordData := utils.StringsFromRegex(in, `^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$`)

	valid := 0
	for _, p := range passwordData {
		first := utils.MustAtoi(p[1]) - 1
		second := utils.MustAtoi(p[2]) - 1
		char := p[3][0]
		password := p[4]

		isFirst := password[first] == char
		isSecond := password[second] == char
		if isFirst == !isSecond {
			valid++
		}
	}
	return strconv.Itoa(valid)

}

func day2StringToCharCounts(in string) map[rune]int {
	counts := map[rune]int{}
	for _, c := range in {
		counts[c]++
	}
	return counts
}
