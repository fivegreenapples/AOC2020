package days

import "fmt"

func d22NewGameOfRecursiveCombat(p1, p2 []int) *d22RecursiveCombat {
	game := d22RecursiveCombat{
		hands: [2]d22PlayerHand{
			{
				cards:     make([]int, len(p1)+len(p2)+1),
				takeIdx:   0,
				insertIdx: len(p1),
			},
			{
				cards:     make([]int, len(p1)+len(p2)+1),
				takeIdx:   0,
				insertIdx: len(p2),
			},
		},
		seenHands: map[[51]int]bool{},
		recurse:   true,
	}

	copy(game.hands[0].cards, p1)
	copy(game.hands[1].cards, p2)
	return &game
}
func d22NewGameOfCombat(p1, p2 []int) *d22RecursiveCombat {
	game := d22RecursiveCombat{
		hands: [2]d22PlayerHand{
			{
				cards:     make([]int, len(p1)+len(p2)+1),
				takeIdx:   0,
				insertIdx: len(p1),
			},
			{
				cards:     make([]int, len(p1)+len(p2)+1),
				takeIdx:   0,
				insertIdx: len(p2),
			},
		},
		seenHands: map[[51]int]bool{},
		recurse:   false,
	}

	copy(game.hands[0].cards, p1)
	copy(game.hands[1].cards, p2)
	return &game
}

type d22PlayerHand struct {
	cards     []int
	takeIdx   int
	insertIdx int
}

type d22RecursiveCombat struct {
	hands     [2]d22PlayerHand
	seenHands map[[51]int]bool
	recurse   bool
}

func (c *d22RecursiveCombat) playToCompletion() int {
	for !c.isWon() {
		handHash := c.hash()
		if c.seenHands[handHash] {
			return 0
		}
		c.seenHands[handHash] = true
		c.playRound()
	}
	return c.winnerIdx()
}

func (c *d22RecursiveCombat) hash() [51]int {
	out := [51]int{}

	offset := 0
	for i := range []int{0, 1} {

		for idx := c.hands[i].takeIdx; idx < c.hands[i].insertIdx; idx++ {
			out[offset] = c.hands[i].cards[idx%len(c.hands[i].cards)]
			offset++
		}
		offset++ // leave one integer gap between hands
	}

	return out
}

func (c *d22RecursiveCombat) playRound() {

	var winnerIdx int

	if c.recurse && c.hands[0].canRecurse() && c.hands[1].canRecurse() {
		// recurse
		subStacks := [2][]int{}
		for i := range []int{0, 1} {
			subStacks[i] = []int{}
			startIdx := c.hands[i].takeIdx + 1
			for idx := startIdx; idx < startIdx+c.hands[i].cards[c.hands[i].takeIdx%len(c.hands[i].cards)]; idx++ {
				subStacks[i] = append(subStacks[i], c.hands[i].cards[idx%len(c.hands[i].cards)])
			}
		}
		subGame := d22NewGameOfRecursiveCombat(subStacks[0], subStacks[1])
		winnerIdx = subGame.playToCompletion()

	} else if c.hands[0].cards[c.hands[0].takeIdx%len(c.hands[0].cards)] > c.hands[1].cards[c.hands[1].takeIdx%len(c.hands[1].cards)] {

		winnerIdx = 0

	} else {

		winnerIdx = 1

	}

	loserIdx := (winnerIdx + 1) % 2

	winner := &c.hands[winnerIdx]
	loser := &c.hands[loserIdx]

	winner.cards[winner.insertIdx%len(winner.cards)] = winner.cards[winner.takeIdx%len(winner.cards)]
	winner.insertIdx++
	winner.cards[winner.insertIdx%len(winner.cards)] = loser.cards[loser.takeIdx%len(loser.cards)]
	winner.insertIdx++

	c.hands[0].takeIdx++
	c.hands[1].takeIdx++
}

func (c *d22RecursiveCombat) isWon() bool {
	return c.hands[0].hasLost() || c.hands[1].hasLost()
}

func (c *d22RecursiveCombat) winnerIdx() int {
	if c.hands[0].hasLost() {
		return 1
	}
	return 0
}

func (c *d22RecursiveCombat) winnersScore() int {
	return c.hands[c.winnerIdx()].score()
}

func (c *d22RecursiveCombat) String() string {
	str := ""
	str += fmt.Sprintf("p1: %s\n", c.hands[0])
	str += fmt.Sprintf("p2: %s\n", c.hands[1])
	return str
}

func (h d22PlayerHand) canRecurse() bool {
	stackLength := h.insertIdx - h.takeIdx
	return (stackLength - 1) >= h.cards[h.takeIdx%len(h.cards)]
}

func (h d22PlayerHand) hasLost() bool {
	return h.takeIdx == h.insertIdx
}

func (h d22PlayerHand) score() int {
	multiplier := 1
	score := 0
	for idx := h.insertIdx - 1; idx >= h.takeIdx; idx-- {
		score += h.cards[idx%len(h.cards)] * multiplier
		multiplier++
	}
	return score
}

func (h d22PlayerHand) String() string {
	out := ""
	if h.takeIdx == h.insertIdx {
		return out
	}

	for idx := h.takeIdx; idx < h.insertIdx; idx++ {
		out += fmt.Sprintf("%d ", h.cards[idx%len(h.cards)])
	}

	return out
}
