package days

import (
	"fmt"
	"strconv"
)

func (r *Runner) Day23Part1(in string) string {
	lookup := r.d23ProcessInputAndPlay(in, 9, 100)
	return r.d23GetValue(lookup)
}

func (r *Runner) Day23Part2(in string) string {
	lookup := r.d23ProcessInputAndPlay(in, 1000000, 10000000)
	return strconv.Itoa(lookup[1] * lookup[lookup[1]])
}

func (r *Runner) d23ProcessInputAndPlay(in string, numCups int, numRounds int) []int {

	lookup := make([]int, numCups+1)

	firstVal := int(in[0] - '0')
	prevVal := firstVal
	for _, rn := range in[1:] {
		lookup[prevVal] = int(rn - '0')
		prevVal = int(rn - '0')
	}
	for i := len(in) + 1; i <= numCups; i++ {
		lookup[prevVal] = i
		prevVal = i
	}
	lookup[prevVal] = firstVal

	var destinationVal int
	val := firstVal
	for i := 0; i < numRounds; i++ {

		/*
			The crab picks up the three cups that are immediately clockwise of the
			current cup. They are removed from the circle; cup spacing is adjusted as
			necessary to maintain the circle.

			The crab selects a destination cup: the cup with a label equal to the
			current cup's label minus one. If this would select one of the cups that was
			just picked up, the crab will keep subtracting one until it finds a cup that
			wasn't just picked up. If at any point in this process the value goes below
			the lowest value on any cup's label, it wraps around to the highest value on
			any cup's label instead.

			The crab places the cups it just picked up so that they are immediately
			clockwise of the destination cup. They keep the same order as when they were
			picked up.

			The crab selects a new current cup: the cup which is immediately clockwise
			of the current cup.
		*/

		aVal := lookup[val]
		bVal := lookup[aVal]
		cVal := lookup[bVal]
		dVal := lookup[cVal]
		// find destination
		destinationVal = val - 1
		for destinationVal == 0 ||
			destinationVal == aVal ||
			destinationVal == bVal ||
			destinationVal == cVal {
			if destinationVal == 0 {
				destinationVal = numCups
			} else {
				destinationVal--
			}
		}

		// connect current to one after the group of three
		lookup[val] = dVal

		// insert group of three
		lookup[cVal] = lookup[destinationVal]
		lookup[destinationVal] = aVal

		val = dVal
	}

	return lookup
}

// }

// func (c d23Cup) String() string {

// curVal := c.val
// str := fmt.Sprint(c.val)
// c = *c.next
// for c.val != curVal {
// 	str += fmt.Sprint(c.val)
// 	c = *c.next
// }

// return str
// }

func (r *Runner) d23GetValue(lookup []int) string {

	c := lookup[1]

	str := ""
	for c != 1 {
		str += fmt.Sprint(c)
		c = lookup[c]
	}

	return str
}
