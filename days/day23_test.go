package days

import "testing"

func TestDay23Part1(t *testing.T) {

	testInputs := map[string]string{
		`389125467`: "67384529",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(23, 1, in)
		if out != expectedOut {
			t.Errorf("day23 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay23Part2(t *testing.T) {

	testInputs := map[string]string{
		`389125467`: "149245887792",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(23, 2, in)
		if out != expectedOut {
			t.Errorf("day23 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
