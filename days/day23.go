package days

import (
	"fmt"
	"strconv"
)

func (r *Runner) Day23Part1(in string) string {
	oneCup := r.d23ProcessInputAndPlay(in, 9, 100)
	return oneCup.value()
}

func (r *Runner) Day23Part2(in string) string {
	oneCup := r.d23ProcessInputAndPlay(in, 1000000, 10000000)
	return strconv.Itoa(oneCup.next.val * oneCup.next.next.val)
}

func (r *Runner) d23ProcessInputAndPlay(in string, numCups int, numRounds int) *d23Cup {

	var firstCup d23Cup
	lookup := map[int]*d23Cup{}

	firstCup = d23Cup{
		val: int(in[0] - '0'),
	}
	lookup[firstCup.val] = &firstCup

	thisCup := &firstCup
	for _, rn := range in[1:] {
		thisCup.next = &d23Cup{
			val: int(rn - '0'),
		}
		lookup[thisCup.next.val] = thisCup.next
		thisCup = thisCup.next
	}
	for i := len(in) + 1; i <= numCups; i++ {
		thisCup.next = &d23Cup{
			val: i,
		}
		lookup[thisCup.next.val] = thisCup.next
		thisCup = thisCup.next
	}
	thisCup.next = &firstCup

	thisCup = &firstCup
	for i := 0; i < numRounds; i++ {
		thisCup.playRound(lookup, numCups)
		thisCup = thisCup.next
	}

	return lookup[1]
}

type d23Cup struct {
	val  int
	next *d23Cup
}

func (c *d23Cup) playRound(lookup map[int]*d23Cup, max int) {

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

	// find destination
	destinationVal := c.val - 1
	for destinationVal == 0 ||
		destinationVal == c.next.val ||
		destinationVal == c.next.next.val ||
		destinationVal == c.next.next.next.val {
		if destinationVal == 0 {
			destinationVal = max
		} else {
			destinationVal--
		}
	}
	destination := lookup[destinationVal]

	// grab reference to head and tail of group of three
	firstOfThree := c.next
	lastOfThree := c.next.next.next

	// connect current to one after the group of three
	c.next = c.next.next.next.next

	// insert group of three
	lastOfThree.next = destination.next
	destination.next = firstOfThree

}

func (c d23Cup) String() string {

	curVal := c.val
	str := fmt.Sprint(c.val)
	c = *c.next
	for c.val != curVal {
		str += fmt.Sprint(c.val)
		c = *c.next
	}

	return str
}

func (c d23Cup) value() string {

	for c.val != 1 {
		c = *c.next
	}
	c = *c.next

	str := ""
	for c.val != 1 {
		str += fmt.Sprint(c.val)
		c = *c.next
	}

	return str
}
