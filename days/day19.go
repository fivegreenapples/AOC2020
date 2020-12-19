package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

type d19RuleType string

const (
	val         d19RuleType = "_val"
	sequence                = "_seq"
	sequenceAlt             = "_seqAlt"
)

type d19Rule struct {
	typ        d19RuleType
	val        string
	ruleSeq    []int
	ruleSeqAlt []int
}

func (r *Runner) Day19Part1(in string) string {

	allRules, messages := r.d19AnalyseInput(in)

	validMessageCount := 0
	for _, m := range messages {

		result := r.d19ConsumeRule(m, 0, allRules)
		if len(result) == 0 {
			continue
		}
		for _, suff := range result {
			if suff == "" {
				if r.verbose {
					fmt.Println("Matched:", m[1])
				}
				validMessageCount++
			}
		}

	}

	return strconv.Itoa(validMessageCount)
}

func (r *Runner) Day19Part2(in string) string {

	allRules, messages := r.d19AnalyseInput(in)

	allRules[8] = d19Rule{
		typ:        sequenceAlt,
		ruleSeq:    []int{42},
		ruleSeqAlt: []int{42, 8},
	}
	allRules[11] = d19Rule{
		typ:        sequenceAlt,
		ruleSeq:    []int{42, 31},
		ruleSeqAlt: []int{42, 11, 31},
	}

	validMessageCount := 0
	for _, m := range messages {

		result := r.d19ConsumeRule(m, 0, allRules)
		if len(result) == 0 {
			continue
		}
		for _, suff := range result {
			if suff == "" {
				if r.verbose {
					fmt.Println("Matched:", m[1])
				}
				validMessageCount++
			}
		}

	}
	return strconv.Itoa(validMessageCount)
}

func (r *Runner) d19AnalyseInput(in string) (map[int]d19Rule, []string) {
	baseRules := utils.StringsFromRegex(in, `^([0-9]+): "([ab])"$`)
	rules := utils.StringsFromRegex(in, `^([0-9]+): ([0-9 |]+)$`)
	messages := utils.StringsFromRegex(in, `^([ab]+)$`)

	allRules := map[int]d19Rule{}

	for _, r := range baseRules {
		if r == nil {
			continue
		}
		allRules[utils.MustAtoi(r[1])] = d19Rule{
			typ: val,
			val: r[2],
		}
	}

	for _, r := range rules {
		if r == nil {
			continue
		}

		thisRule := d19Rule{}

		splitAlts := strings.Split(r[2], "|")
		if len(splitAlts) == 1 {
			thisRule.typ = sequence
		} else {
			thisRule.typ = sequenceAlt
		}
		for splitIdx, alt := range splitAlts {
			rules := strings.Split(strings.TrimSpace(alt), " ")
			if splitIdx == 0 {
				for _, r := range rules {
					thisRule.ruleSeq = append(thisRule.ruleSeq, utils.MustAtoi(r))
				}
			} else {
				for _, r := range rules {
					thisRule.ruleSeqAlt = append(thisRule.ruleSeqAlt, utils.MustAtoi(r))
				}
			}
		}

		allRules[utils.MustAtoi(r[1])] = thisRule
	}

	validMessages := []string{}
	for _, m := range messages {
		if m == nil {
			continue
		}
		validMessages = append(validMessages, m[1])
	}

	return allRules, validMessages
}

func (r *Runner) d19ConsumeRule(in string, ruleNum int, allRules map[int]d19Rule) []string {

	rule := allRules[ruleNum]

	if rule.typ == val {
		if strings.HasPrefix(in, rule.val) {
			return []string{
				strings.TrimPrefix(in, rule.val),
			}
		}
		return nil
	}
	if rule.typ == sequence {
		availableInputs := []string{in}
		for _, subRuleNum := range rule.ruleSeq {
			thisOutputs := []string{}
			for _, input := range availableInputs {
				thisOutputs = append(thisOutputs, r.d19ConsumeRule(input, subRuleNum, allRules)...)
			}
			availableInputs = thisOutputs
		}
		return availableInputs
	}

	if rule.typ == sequenceAlt {
		availableInputs := []string{in}
		for _, subRuleNum := range rule.ruleSeq {
			thisOutputs := []string{}
			for _, input := range availableInputs {
				thisOutputs = append(thisOutputs, r.d19ConsumeRule(input, subRuleNum, allRules)...)
			}
			availableInputs = thisOutputs
		}

		availableInputsAlt := []string{in}
		for _, subRuleNum := range rule.ruleSeqAlt {
			thisOutputs := []string{}
			for _, input := range availableInputsAlt {
				thisOutputs = append(thisOutputs, r.d19ConsumeRule(input, subRuleNum, allRules)...)
			}
			availableInputsAlt = thisOutputs
		}

		return append(availableInputs, availableInputsAlt...)
	}

	return nil
}
