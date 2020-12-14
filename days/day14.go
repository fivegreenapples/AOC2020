package days

import (
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day14Part1(in string) string {

	instructions := utils.StringsFromRegex(in, `^((mask) = ([X10]+))|((mem)\[([0-9]+)\] = ([0-9]+))$`)

	mem := map[int]int{}
	var andMask int
	var orMask int
	for _, instr := range instructions {

		if instr[2] == "mask" {
			mask := instr[3]
			andMask = utils.MustParseBinary(strings.ReplaceAll(mask, "X", "1"))
			orMask = utils.MustParseBinary(strings.ReplaceAll(mask, "X", "0"))
		} else {

			addr := utils.MustAtoi(instr[6])
			value := utils.MustAtoi(instr[7])

			value = value&andMask | orMask
			mem[addr] = int(value)
		}

	}

	sum := 0
	for _, val := range mem {
		sum += val
	}
	return strconv.Itoa(sum)
}

func (r *Runner) Day14Part2(in string) string {

	instructions := utils.StringsFromRegex(in, `^((mask) = ([X10]+))|((mem)\[([0-9]+)\] = ([0-9]+))$`)

	mem := map[int]int{}
	var mask string
	for _, instr := range instructions {

		if instr[2] == "mask" {

			mask = instr[3]

		} else {
			addr := utils.MustAtoi(instr[6])
			value := utils.MustAtoi(instr[7])

			// clear any bits in address at 'X' positions in mask
			andMask := utils.MustParseBinary(
				strings.ReplaceAll(
					strings.ReplaceAll(mask, "0", "1"),
					"X",
					"0",
				),
			)
			addr = addr & andMask

			addressMasks := r.day14GetAddressMasks(mask)
			for _, orMask := range addressMasks {
				thisAddr := addr | orMask
				mem[thisAddr] = int(value)
			}

		}

	}

	sum := 0
	for _, val := range mem {
		sum += val
	}
	return strconv.Itoa(sum)
}

func (r *Runner) day14GetAddressMasks(mask string) []int {

	if strings.Count(mask, "X") == 0 {
		return []int{utils.MustParseBinary(mask)}
	}

	vals := []int{}
	for _, replacement := range []string{"0", "1"} {
		newMask := strings.Replace(mask, "X", replacement, 1)
		vals = append(vals, r.day14GetAddressMasks(newMask)...)
	}
	return vals
}
