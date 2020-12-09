package days

import (
	"testing"

	"github.com/fivegreenapples/AOC2020/utils"
)

func TestDay9Part1(t *testing.T) {

	testInputs := map[string]int{
		`35
		20
		15
		25
		47
		40
		62
		55
		65
		95
		102
		117
		150
		182
		127
		219
		299
		277
		309
		576`: 127,
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out := dayRunner.day9FirstInvalid(5, utils.LinesAsInts(in))
		if out != expectedOut {
			t.Errorf("day9 pt1 failed with %s. Expected %d, got %d", in, expectedOut, out)
		}
	}

}

func TestDay9Part2(t *testing.T) {

	testInputs := map[string]int{
		`35
		20
		15
		25
		47
		40
		62
		55
		65
		95
		102
		117
		150
		182
		127
		219
		299
		277
		309
		576`: 62,
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out := dayRunner.day9EncryptionWeakness(5, utils.LinesAsInts(in))
		if out != expectedOut {
			t.Errorf("day9 pt1 failed with %s. Expected %d, got %d", in, expectedOut, out)
		}
	}

}
