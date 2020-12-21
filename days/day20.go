package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

type d20Tile struct {
	id          int
	top         [10]byte
	right       [10]byte
	bottom      [10]byte
	left        [10]byte
	grid        [10][10]byte
	isCorner    bool
	topMatch    int
	rightMatch  int
	bottomMatch int
	leftMatch   int
}

func (t *d20Tile) calcSides() {
	t.top = t.grid[0]
	t.bottom = t.grid[9]
	for r := 0; r < 10; r++ {
		t.left[r] = t.grid[r][0]
		t.right[r] = t.grid[r][9]
	}
}

func (t *d20Tile) calcMatchAgainst(tt d20Tile) {
	numMatches := 0
	if t.top == tt.top || t.top == tt.right || t.top == tt.bottom || t.top == tt.left {
		t.topMatch = tt.id
		numMatches++
	}
	if t.right == tt.top || t.right == tt.right || t.right == tt.bottom || t.right == tt.left {
		t.rightMatch = tt.id
		numMatches++
	}
	if t.bottom == tt.top || t.bottom == tt.right || t.bottom == tt.bottom || t.bottom == tt.left {
		t.bottomMatch = tt.id
		numMatches++
	}
	if t.left == tt.top || t.left == tt.right || t.left == tt.bottom || t.left == tt.left {
		t.leftMatch = tt.id
		numMatches++
	}
	topFlipped := d20FlipSide(t.top)
	rightFlipped := d20FlipSide(t.right)
	bottomFlipped := d20FlipSide(t.bottom)
	leftFlipped := d20FlipSide(t.left)
	if topFlipped == tt.top || topFlipped == tt.right || topFlipped == tt.bottom || topFlipped == tt.left {
		t.topMatch = tt.id
		numMatches++
	}
	if rightFlipped == tt.top || rightFlipped == tt.right || rightFlipped == tt.bottom || rightFlipped == tt.left {
		t.rightMatch = tt.id
		numMatches++
	}
	if bottomFlipped == tt.top || bottomFlipped == tt.right || bottomFlipped == tt.bottom || bottomFlipped == tt.left {
		t.bottomMatch = tt.id
		numMatches++
	}
	if leftFlipped == tt.top || leftFlipped == tt.right || leftFlipped == tt.bottom || leftFlipped == tt.left {
		t.leftMatch = tt.id
		numMatches++
	}
	if numMatches > 1 {
		panic(fmt.Errorf("too many matches for tile %d against %d", t.id, tt.id))
	}

	t.calcCornerStatus()
}

func (t *d20Tile) calcCornerStatus() {
	numMatches := 0
	if t.topMatch > 0 {
		numMatches++
	}
	if t.rightMatch > 0 {
		numMatches++
	}
	if t.bottomMatch > 0 {
		numMatches++
	}
	if t.leftMatch > 0 {
		numMatches++
	}
	t.isCorner = numMatches == 2
}

func (t *d20Tile) rotateClockwise() {
	top, right, bottom, left := d20FlipSide(t.left), t.top, d20FlipSide(t.right), t.bottom
	t.topMatch, t.rightMatch, t.bottomMatch, t.leftMatch = t.leftMatch, t.topMatch, t.rightMatch, t.bottomMatch
	newGrid := [10][10]byte{}
	for c := 0; c < 10; c++ {
		for r := 9; r >= 0; r-- {
			newGrid[c][9-r] = t.grid[r][c]
		}
	}
	t.grid = newGrid
	t.calcSides()
	if t.top != top || t.right != right || t.bottom != bottom || t.left != left {
		panic(fmt.Errorf("rotate error with %d", t.id))
	}

}
func (t *d20Tile) flipTopBottom() {
	top, right, bottom, left := t.bottom, d20FlipSide(t.right), t.top, d20FlipSide(t.left)
	t.topMatch, t.bottomMatch = t.bottomMatch, t.topMatch
	newGrid := [10][10]byte{}
	for r := 9; r >= 0; r-- {
		newGrid[9-r] = t.grid[r]
	}
	t.grid = newGrid
	t.calcSides()
	if t.top != top || t.right != right || t.bottom != bottom || t.left != left {
		panic(fmt.Errorf("flip vertical error with %d", t.id))
	}

}

func (r *Runner) Day20Part1(in string) string {

	lines := utils.Lines(in)
	tiles := []d20Tile{}

	for l := 0; l < len(lines); l++ {

		thisLine := lines[l]
		if !strings.HasPrefix(thisLine, "Tile ") {
			fmt.Println("error parsing input")
			return "?"
		}

		thisTile := d20Tile{
			id: utils.MustAtoi(thisLine[5 : len(thisLine)-1]),
		}

		for r := 0; r < 10; r++ {
			l++
			thisLine = lines[l]
			for c := 0; c < 10; c++ {
				thisTile.grid[r][c] = thisLine[c]
			}
		}
		thisTile.calcSides()
		l++
		tiles = append(tiles, thisTile)

	}

	for x := range tiles {
		for y, tY := range tiles {
			if x == y {
				continue
			}
			tiles[x].calcMatchAgainst(tY)
		}
	}

	cornerProduct := 1
	for _, t := range tiles {
		if t.isCorner {
			cornerProduct *= t.id
		}
	}

	return strconv.Itoa(cornerProduct)
}

func (r *Runner) Day20Part2(in string) string {

	lines := utils.Lines(in)
	tiles := []d20Tile{}

	for l := 0; l < len(lines); l++ {

		thisLine := lines[l]
		if !strings.HasPrefix(thisLine, "Tile ") {
			fmt.Println("error parsing input")
			return "?"
		}

		thisTile := d20Tile{
			id: utils.MustAtoi(thisLine[5 : len(thisLine)-1]),
		}

		for r := 0; r < 10; r++ {
			l++
			thisLine = lines[l]
			for c := 0; c < 10; c++ {
				thisTile.grid[r][c] = thisLine[c]
			}
		}
		thisTile.calcSides()
		l++
		tiles = append(tiles, thisTile)

	}

	for x := range tiles {
		for y, tY := range tiles {
			if x == y {
				continue
			}
			tiles[x].calcMatchAgainst(tY)
		}
	}

	// Find a corner, and create tile map
	firstCornerId := 0
	tileMap := map[int]*d20Tile{}
	for idx, t := range tiles {
		if t.isCorner {
			firstCornerId = t.id
		}
		tileMap[t.id] = &tiles[idx]
	}

	// orientate first corner to top left
	for tileMap[firstCornerId].leftMatch > 0 {
		tileMap[firstCornerId].rotateClockwise()
	}
	if tileMap[firstCornerId].topMatch > 0 {
		tileMap[firstCornerId].rotateClockwise()
	}

	sideLength := int(math.Sqrt(float64(len(tiles))))

	finishedGrid := make([][]int, sideLength)
	for row := 0; row < sideLength; row++ {
		currentRow := make([]int, sideLength)

		if row == 0 {
			currentRow[0] = firstCornerId
		} else {
			tileAbove := tileMap[finishedGrid[row-1][0]]
			newEdgeTile := tileMap[tileAbove.bottomMatch]
			if newEdgeTile.topMatch == tileAbove.id {
			} else if newEdgeTile.leftMatch == tileAbove.id {
				newEdgeTile.rotateClockwise()
			} else if newEdgeTile.bottomMatch == tileAbove.id {
				newEdgeTile.rotateClockwise()
				newEdgeTile.rotateClockwise()
			} else if newEdgeTile.rightMatch == tileAbove.id {
				newEdgeTile.rotateClockwise()
				newEdgeTile.rotateClockwise()
				newEdgeTile.rotateClockwise()
			} else {
				panic("edge failure in grid assembly - new tile doesn't match against current")
			}
			if newEdgeTile.top != tileAbove.bottom {
				newEdgeTile.flipTopBottom()
				newEdgeTile.rotateClockwise()
				newEdgeTile.rotateClockwise()
			}
			if newEdgeTile.top != tileAbove.bottom {
				panic("edge failure in grid assembly after side alignment")
			}
			currentRow[0] = newEdgeTile.id

		}
		for col := 1; col < sideLength; col++ {
			tile := tileMap[currentRow[col-1]]
			nextTile := tileMap[tile.rightMatch]
			if nextTile.leftMatch == tile.id {
			} else if nextTile.bottomMatch == tile.id {
				nextTile.rotateClockwise()
			} else if nextTile.rightMatch == tile.id {
				nextTile.rotateClockwise()
				nextTile.rotateClockwise()
			} else if nextTile.topMatch == tile.id {
				nextTile.rotateClockwise()
				nextTile.rotateClockwise()
				nextTile.rotateClockwise()
			} else {
				panic("failure in grid assembly - new tile doesn't match against current")
			}
			if nextTile.left != tile.right {
				nextTile.flipTopBottom()
			}
			if nextTile.left != tile.right {
				panic("failure in grid assembly after side alignment")
			}
			currentRow[col] = nextTile.id
		}
		finishedGrid[row] = currentRow
	}

	// construct borderless full grid
	pixelLength := sideLength * 8
	pixels := make([][]byte, pixelLength)

	for rowIdx, row := range finishedGrid {
		pixels[(rowIdx * 8)] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+1] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+2] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+3] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+4] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+5] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+6] = make([]byte, pixelLength)
		pixels[(rowIdx*8)+7] = make([]byte, pixelLength)
		for colIdx, tileId := range row {

			tile := tileMap[tileId]

			for rIdx := 1; rIdx <= 8; rIdx++ {
				for cIdx := 1; cIdx <= 8; cIdx++ {

					pixels[(rowIdx*8)+rIdx-1][(colIdx*8)+cIdx-1] = tile.grid[rIdx][cIdx]

				}

			}

		}

	}

	if r.verbose {
		d20PrintGrid(pixels)
	}

	monster := [][]byte{
		[]byte(`                  # `),
		[]byte(`#    ##    ##    ###`),
		[]byte(` #  #  #  #  #  #   `),
	}

	// Try all 8 orientations to find monsters. Stop as soon as we have at least one monster.
	numMonsters := 0
	for r := 0; r < 4; r++ {
		numMonsters = d20CountPatternInGrid(pixels, monster)
		if numMonsters > 0 {
			break
		}
		pixels = d20RotateGrid(pixels)
	}
	if numMonsters == 0 {
		pixels = d20FlipGrid(pixels)
		for r := 0; r < 4; r++ {
			numMonsters = d20CountPatternInGrid(pixels, monster)
			if numMonsters > 0 {
				break
			}
			pixels = d20RotateGrid(pixels)
		}
	}
	if numMonsters == 0 {
		return "no monsters found :("
	}

	if r.verbose {
		d20PrintGridWithPattern(pixels, monster)
	}

	numWaves := d20CountByteInGrid(pixels, '#') - numMonsters*d20CountByteInGrid(monster, '#')
	return strconv.Itoa(numWaves)
}

func d20FlipSide(in [10]byte) [10]byte {
	for left, right := 0, len(in)-1; left < right; left, right = left+1, right-1 {
		in[left], in[right] = in[right], in[left]
	}
	return in
}
func d20CountByteInGrid(grid [][]byte, pattern byte) int {
	count := 0
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[rowIdx]); colIdx++ {
			if grid[rowIdx][colIdx] == pattern {
				count++
			}
		}
	}
	return count
}

func d20CountPatternInGrid(grid [][]byte, pattern [][]byte) int {

	// This assumes the patterns don't overlap
	count := 0
	for rowIdx := 0; rowIdx < len(grid)-len(pattern)+1; rowIdx++ {
		lineIndices := d20IndexAll(grid[rowIdx], pattern[0])
		for p := 1; p < len(pattern); p++ {
			newLineIndices := []int{}
			for _, idx := range lineIndices {
				if d20HasPatternPrefix(grid[rowIdx+p][idx:], pattern[p]) {
					newLineIndices = append(newLineIndices, idx)
				}
			}
			lineIndices = newLineIndices
			if len(lineIndices) == 0 {
				break
			}
		}
		count += len(lineIndices)
	}

	return count
}
func d20IndexAll(line []byte, pattern []byte) []int {
	indices := []int{}
	for l := 0; l < len(line)-len(pattern)+1; l++ {
		matched := true
		for p := range pattern {
			if pattern[p] != '#' {
				continue
			}
			if line[l+p] != '#' {
				matched = false
				break
			}
		}
		if matched {
			indices = append(indices, l)
		}
	}
	return indices
}

func d20HasPatternPrefix(line []byte, pattern []byte) bool {
	if len(pattern) > len(line) {
		return false
	}
	matched := true
	for p := 0; p < len(pattern); p++ {
		if pattern[p] != '#' {
			continue
		}
		if line[p] != '#' {
			matched = false
			break
		}
	}
	return matched
}
func d20FlipGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for r := len(grid) - 1; r >= 0; r-- {
		newGrid[len(grid)-1-r] = grid[r]
	}
	return newGrid
}
func d20RotateGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for c := 0; c < len(grid); c++ {
		newGrid[c] = make([]byte, len(grid))
		for r := len(grid) - 1; r >= 0; r-- {
			newGrid[c][len(grid)-1-r] = grid[r][c]
		}
	}
	return newGrid
}

func d20PrintGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func d20PrintGridWithPattern(grid, pattern [][]byte) {
	const (
		PURPLE = "\033[0;35m"
		CLEAR  = "\033[0m"
	)
	for rowIdx := 0; rowIdx < len(grid)-len(pattern)+1; rowIdx++ {
		lineIndices := d20IndexAll(grid[rowIdx], pattern[0])
		for p := 1; p < len(pattern); p++ {
			newLineIndices := []int{}
			for _, idx := range lineIndices {
				if d20HasPatternPrefix(grid[rowIdx+p][idx:], pattern[p]) {
					newLineIndices = append(newLineIndices, idx)
				}
			}
			lineIndices = newLineIndices
			if len(lineIndices) == 0 {
				break
			}
		}

		if len(lineIndices) > 0 {
			for lineOffset := 0; lineOffset < len(pattern); lineOffset++ {
				thisLine := make([]byte, len(grid[rowIdx]))
				for chIdx, ch := range grid[rowIdx] {
					thisLine[chIdx] = ch
				}
				for _, lineIdx := range lineIndices {
					for p := 0; p < len(pattern[lineOffset]); p++ {
						if pattern[lineOffset][p] == '#' {
							thisLine[lineIdx+p] = 'O'
						}
					}
				}
				colourfulString := ""
				for _, b := range thisLine {
					if b == 'O' {
						colourfulString += PURPLE + "O" + CLEAR
					} else {
						colourfulString += string(b)
					}
				}
				fmt.Println(colourfulString)
			}
			rowIdx += len(pattern) - 1
		} else {
			fmt.Println(string(grid[rowIdx]))
		}
	}
	for rowIdx := len(grid) - len(pattern) + 1; rowIdx < len(grid); rowIdx++ {
		fmt.Println(string(grid[rowIdx]))
	}

	fmt.Println()
}
