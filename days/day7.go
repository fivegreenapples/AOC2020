package days

import (
	"fmt"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

type day7BagRuleset map[string]int

func (r *Runner) Day7Part1(in string) string {

	ruleMatches := utils.StringsFromRegex(in, `^(.+?) bags contain (([0-9]+ .+?)|(no other bags))\.$`)

	rules := map[string]day7BagRuleset{}

	for _, ruleMatch := range ruleMatches {

		thisRuleSet := day7BagRuleset{}

		thisBag := ruleMatch[1]
		thisRuleDescription := ruleMatch[2]
		if thisRuleDescription == "no other bags" {
			// nothing to do
			if r.verbose {
				fmt.Printf("%s - no other bags\n", thisBag)
			}
		} else {
			if r.verbose {
				fmt.Printf("%s - has bags:\n", thisBag)
			}

			descriptionMatches := utils.AllStringsFromRegex(thisRuleDescription, `([0-9]+) ([^,]+) bag`)
			for _, d := range descriptionMatches[0] {

				thisCount := utils.MustAtoi(d[1])
				thisBag := d[2]
				if r.verbose {
					fmt.Printf(" - %d : %s\n", thisCount, thisBag)
				}

				thisRuleSet[thisBag] = thisCount

			}

		}

		rules[thisBag] = thisRuleSet
	}

	// Create reverse rules - i.e. bags which can be contained by
	reverseRules := map[string]day7BagRuleset{}
	for containingBag, ruleset := range rules {
		for containedBag := range ruleset {
			currentruleset := reverseRules[containedBag]
			if currentruleset == nil {
				currentruleset = day7BagRuleset{}
			}
			currentruleset[containingBag] = 1

			reverseRules[containedBag] = currentruleset
		}
	}

	eventuallyContainingBags := map[string]bool{}

	testBags := []string{"shiny gold"}
	for len(testBags) != 0 {
		thisTestBag := testBags[len(testBags)-1]
		testBags = testBags[:len(testBags)-1]

		for b := range reverseRules[thisTestBag] {
			if eventuallyContainingBags[b] {
				continue
			}
			eventuallyContainingBags[b] = true
			testBags = append(testBags, b)
		}

	}

	if r.verbose {
		fmt.Println(eventuallyContainingBags)
	}

	return strconv.Itoa(len(eventuallyContainingBags))
}

func (r *Runner) Day7Part2(in string) string {

	ruleMatches := utils.StringsFromRegex(in, `^(.+?) bags contain (([0-9]+ .+?)|(no other bags))\.$`)

	rules := map[string]day7BagRuleset{}
	emptyBags := []string{}

	for _, ruleMatch := range ruleMatches {

		thisRuleSet := day7BagRuleset{}

		thisBag := ruleMatch[1]
		thisRuleDescription := ruleMatch[2]
		if thisRuleDescription == "no other bags" {
			if r.verbose {
				fmt.Printf("%s - no other bags\n", thisBag)
			}
			emptyBags = append(emptyBags, thisBag)
		} else {
			if r.verbose {
				fmt.Printf("%s - has bags:\n", thisBag)
			}

			descriptionMatches := utils.AllStringsFromRegex(thisRuleDescription, `([0-9]+) ([^,]+) bag`)
			for _, d := range descriptionMatches[0] {

				thisCount := utils.MustAtoi(d[1])
				thisBag := d[2]
				if r.verbose {
					fmt.Printf(" - %d : %s\n", thisCount, thisBag)
				}

				thisRuleSet[thisBag] = thisCount

			}

		}

		rules[thisBag] = thisRuleSet
	}

	var countForBag func(string) int
	countForBag = func(bag string) int {
		count := 0
		for b, num := range rules[bag] {
			count += num + num*countForBag(b)
		}
		return count
	}

	countForShinyGold := countForBag("shiny gold")

	return strconv.Itoa(countForShinyGold)

}
