package days

import "testing"

func TestDay14Part1(t *testing.T) {

	testInputs := map[string]string{
		`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
		mem[8] = 11
		mem[7] = 101
		mem[8] = 0`: "165",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(14, 1, in)
		if out != expectedOut {
			t.Errorf("day14 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay14Part2(t *testing.T) {

	testInputs := map[string]string{
		`mask = 000000000000000000000000000000X1001X
		mem[42] = 100
		mask = 00000000000000000000000000000000X0XX
		mem[26] = 1`: "208",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(14, 2, in)
		if out != expectedOut {
			t.Errorf("day14 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
