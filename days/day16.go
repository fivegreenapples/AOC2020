package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

type d16Range struct {
	min int
	max int
}
type d16FieldRule struct {
	name   string
	ranges [2]d16Range
}
type d16Ticket []int

func (r d16FieldRule) isValid(v int) bool {
	return (v >= r.ranges[0].min && v <= r.ranges[0].max) || (v >= r.ranges[1].min && v <= r.ranges[1].max)
}

func (r *Runner) Day16Part1(in string) string {

	rangeMatches := utils.StringsFromRegex(in, `^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)
	ticketMatches := utils.StringsFromRegex(in, `^([0-9]+(,[0-9]+)*)$`)

	ticketRules := []d16FieldRule{}
	otherTickets := []d16Ticket{}

	for _, m := range rangeMatches {
		if len(m) == 0 {
			continue
		}
		ticketRules = append(ticketRules, d16FieldRule{
			name: m[1],
			ranges: [2]d16Range{
				{min: utils.MustAtoi(m[2]), max: utils.MustAtoi(m[3])},
				{min: utils.MustAtoi(m[4]), max: utils.MustAtoi(m[5])},
			},
		})
	}

	seenOurTicket := false
	for _, m := range ticketMatches {
		if len(m) == 0 {
			continue
		}

		if !seenOurTicket {
			seenOurTicket = true
			continue
		}

		otherTickets = append(otherTickets, utils.CsvToInts(m[1]))
	}

	errorRate := 0
	for _, t := range otherTickets {

		for _, tVal := range t {

			valid := false
			for _, rule := range ticketRules {
				valid = valid || rule.isValid(tVal)
			}
			if !valid {
				errorRate += tVal
			}

		}
	}

	return strconv.Itoa(errorRate)
}

func (r *Runner) Day16Part2(in string) string {
	rangeMatches := utils.StringsFromRegex(in, `^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)
	ticketMatches := utils.StringsFromRegex(in, `^([0-9]+(,[0-9]+)*)$`)

	ticketRules := []d16FieldRule{}
	otherTickets := []d16Ticket{}
	var ourTicket d16Ticket

	for _, m := range rangeMatches {
		if len(m) == 0 {
			continue
		}
		ticketRules = append(ticketRules, d16FieldRule{
			name: m[1],
			ranges: [2]d16Range{
				{min: utils.MustAtoi(m[2]), max: utils.MustAtoi(m[3])},
				{min: utils.MustAtoi(m[4]), max: utils.MustAtoi(m[5])},
			},
		})
	}

	seenOurTicket := false
	for _, m := range ticketMatches {
		if len(m) == 0 {
			continue
		}

		if !seenOurTicket {
			ourTicket = utils.CsvToInts(m[1])
			seenOurTicket = true
			continue
		}

		otherTickets = append(otherTickets, utils.CsvToInts(m[1]))
	}

	validTickets := []d16Ticket{}
	for _, t := range otherTickets {

		allFieldsHaveOneValidRule := true
		for _, tVal := range t {

			anyRuleValid := false
			for _, rule := range ticketRules {
				anyRuleValid = anyRuleValid || rule.isValid(tVal)
			}
			if !anyRuleValid {
				allFieldsHaveOneValidRule = false
				break
			}
		}
		if allFieldsHaveOneValidRule {
			validTickets = append(validTickets, t)
		}
	}

	fieldMappings := map[string]map[int]bool{}

	for _, rule := range ticketRules {
		for fIdx := 0; fIdx < len(ticketRules); fIdx++ {
			allValid := true
			for _, t := range validTickets {
				tVal := t[fIdx]
				allValid = allValid && rule.isValid(tVal)
			}
			if allValid {
				if fieldMappings[rule.name] == nil {
					fieldMappings[rule.name] = map[int]bool{}
				}
				fieldMappings[rule.name][fIdx] = true
			}
		}
	}

	solvedMappings := map[string]int{}

	for len(solvedMappings) < len(ticketRules) {

		newSolvedField := -1
		for name, availableFields := range fieldMappings {
			if _, exists := solvedMappings[name]; exists {
				continue
			}

			if len(availableFields) == 1 {
				for fieldIdx := range availableFields {
					newSolvedField = fieldIdx
					solvedMappings[name] = newSolvedField
				}
				break
			}
		}

		if newSolvedField == -1 {
			fmt.Println("Failed to solve")
			return "?"
		}

		for name, availableFields := range fieldMappings {
			if _, exists := solvedMappings[name]; exists {
				continue
			}

			delete(availableFields, newSolvedField)
			fieldMappings[name] = availableFields
		}

	}

	departureFieldProduct := 1
	for name, fIdx := range solvedMappings {
		if strings.HasPrefix(name, "departure") {
			departureFieldProduct *= ourTicket[fIdx]
		}
	}
	return strconv.Itoa(departureFieldProduct)
}
