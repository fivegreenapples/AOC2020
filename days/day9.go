package days

import (
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day9Part1(in string) string {
	data := utils.LinesAsInts(in)
	firstInvalid := r.day9FirstInvalid(25, data)
	return strconv.Itoa(firstInvalid)
}

func (r *Runner) Day9Part2(in string) string {
	data := utils.LinesAsInts(in)
	encryptionWeakness := r.day9EncryptionWeakness(25, data)
	return strconv.Itoa(encryptionWeakness)
}

func (r *Runner) day9FirstInvalid(lenPreamble int, data []int) int {

	// loop over data starting immediately after preamble
	for i, val := range data[lenPreamble:] {

		// calculate sums by adding up pairs of numbers from the block of numbers before current index
		start := i
		end := i + lenPreamble
		for a, aVal := range data[start:end] {
			for _, bVal := range data[start+a+1 : end] {
				if aVal+bVal == val {
					goto donesumming
				}
			}
		}

		return val
	donesumming:
	}

	return -1
}

func (r *Runner) day9EncryptionWeakness(lenPreamble int, data []int) int {

	firstInvalid := r.day9FirstInvalid(lenPreamble, data)

	for d := 0; d < len(data); d++ {

		if data[d] >= firstInvalid {
			break
		}

		currentSum := 0
		dd := d
		for currentSum < firstInvalid {
			currentSum += data[dd]
			dd++
		}
		if currentSum == firstInvalid {
			// found match. get slice of nums in sum.
			nums := make([]int, dd-d)
			copy(nums, data[d:dd])
			// then sort
			sort.Ints(nums)
			// then return sum of min and max
			return nums[0] + nums[len(nums)-1]
		}

	}

	return -1
}
