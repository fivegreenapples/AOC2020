package days

import "testing"

func TestDay12Part1(t *testing.T) {

	testInputs := map[string]string{
		`F10
		N3
		F7
		R90
		F11`: "25",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(12, 1, in)
		if out != expectedOut {
			t.Errorf("day12 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay12Part2(t *testing.T) {

	testInputs := map[string]string{
		`F10
		N3
		F7
		R90
		F11`: "286",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(12, 2, in)
		if out != expectedOut {
			t.Errorf("day12 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
