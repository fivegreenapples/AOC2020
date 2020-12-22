package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day22Part1(in string) string {

	players := d22ParseInput(in)
	game := d22NewGameOfCombat(players[0], players[1])
	if r.verbose {
		fmt.Println(game)
	}
	game.playToCompletion()
	return strconv.Itoa(game.winnersScore())
}

func (r *Runner) Day22Part2(in string) string {

	players := d22ParseInput(in)
	game := d22NewGameOfRecursiveCombat(players[0], players[1])
	if r.verbose {
		fmt.Println(game)
	}
	game.playToCompletion()
	return strconv.Itoa(game.winnersScore())
}

func d22ParseInput(in string) [][]int {
	lines := utils.Lines(in)

	players := [][]int{}

	l := 0
	for l < len(lines) {
		line := lines[l]
		if !strings.HasPrefix(line, "Player") {
			panic("failed parse of cards")
		}

		l++
		cards := []int{}
		for l < len(lines) {
			line := lines[l]
			if line == "" {
				break
			}
			cards = append(cards, utils.MustAtoi(line))
			l++
		}
		players = append(players, cards)
		l++
	}

	return players
}
