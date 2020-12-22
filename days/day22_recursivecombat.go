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
		seenHands: map[string]bool{},
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
		seenHands: map[string]bool{},
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
	seenHands map[string]bool
	recurse   bool
}

func (c *d22RecursiveCombat) playToCompletion() int {
	for !c.isWon() {
		handHash := c.hands[0].String() + "|" + c.hands[1].String()
		if c.seenHands[handHash] {
			return 0
		}
		c.seenHands[handHash] = true
		c.playRound()
	}
	return c.winnerIdx()
}
func (c *d22RecursiveCombat) playRound() {

	var winnerIdx int

	if c.recurse && c.hands[0].canRecurse() && c.hands[1].canRecurse() {
		// recurse
		subStacks := [2][]int{}
		for i := range []int{0, 1} {
			subStacks[i] = []int{}
			startIdx := c.hands[i].takeIdx + 1
			for idx := startIdx; idx < startIdx+c.hands[i].cards[c.hands[i].takeIdx]; idx++ {
				subStacks[i] = append(subStacks[i], c.hands[i].cards[idx%len(c.hands[i].cards)])
			}
		}
		subGame := d22NewGameOfRecursiveCombat(subStacks[0], subStacks[1])
		winnerIdx = subGame.playToCompletion()

	} else if c.hands[0].cards[c.hands[0].takeIdx] > c.hands[1].cards[c.hands[1].takeIdx] {

		winnerIdx = 0

	} else {

		winnerIdx = 1

	}

	loserIdx := (winnerIdx + 1) % 2

	winner := &c.hands[winnerIdx]
	loser := &c.hands[loserIdx]

	winner.cards[winner.insertIdx] = winner.cards[winner.takeIdx]
	winner.insertIdx = (winner.insertIdx + 1) % len(winner.cards)
	winner.cards[winner.insertIdx] = loser.cards[loser.takeIdx]
	winner.insertIdx = (winner.insertIdx + 1) % len(winner.cards)

	c.hands[0].takeIdx = (c.hands[0].takeIdx + 1) % len(c.hands[0].cards)
	c.hands[1].takeIdx = (c.hands[1].takeIdx + 1) % len(c.hands[1].cards)
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
	if stackLength <= 0 {
		stackLength += len(h.cards)
	}
	return (stackLength - 1) >= h.cards[h.takeIdx]
}

func (h d22PlayerHand) hasLost() bool {
	return h.takeIdx == h.insertIdx
}

func (h d22PlayerHand) score() int {
	idx := h.insertIdx
	multiplier := 1
	score := 0
	for {
		idx = (idx - 1 + len(h.cards)) % len(h.cards)
		score += h.cards[idx] * multiplier
		multiplier++

		if idx == h.takeIdx {
			break
		}
	}
	return score
}

func (h d22PlayerHand) String() string {
	out := ""
	if h.takeIdx == h.insertIdx {
		return out
	}

	idx := h.takeIdx
	for {
		out += fmt.Sprintf("%d ", h.cards[idx])
		idx = (idx + 1) % len(h.cards)
		if idx == h.insertIdx {
			break
		}
	}

	return out
}
