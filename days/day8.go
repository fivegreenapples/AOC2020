package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day8Part1(in string) string {

	instructions := utils.StringsFromRegex(in, `^([a-z]{3}) ([+-][0-9]+)$`)
	ip := 0
	acc := 0
	seenInstructions := map[int]bool{}
	for {
		if seenInstructions[ip] {
			return strconv.Itoa(acc)
		}
		seenInstructions[ip] = true
		instruction := instructions[ip][1]
		operand := utils.MustAtoi(instructions[ip][2])

		switch instruction {
		case "nop":
			ip++
		case "acc":
			acc += operand
			ip++
		case "jmp":
			ip += operand
		}
	}

}

func (r *Runner) Day8Part2(in string) string {

	instructions := utils.StringsFromRegex(in, `^([a-z]{3}) ([+-][0-9]+)$`)

	for i := 0; i < len(instructions); i++ {

		ip := 0
		acc := 0
		seenInstructions := map[int]bool{}

		for !seenInstructions[ip] {
			if ip >= len(instructions) {
				return strconv.Itoa(acc)
			}
			seenInstructions[ip] = true
			instruction := instructions[ip][1]
			operand := utils.MustAtoi(instructions[ip][2])

			if i == ip {
				if instruction == "jmp" {
					instruction = "nop"
				} else if instruction == "nop" {
					instruction = "jmp"
				}
			}

			switch instruction {
			case "nop":
				ip++
			case "acc":
				acc += operand
				ip++
			case "jmp":
				ip += operand
			}
		}
	}

	return "?"
}
