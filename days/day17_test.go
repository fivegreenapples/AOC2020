package days

import "testing"

func TestDay17Part1(t *testing.T) {

	testInputs := map[string]string{
		`.#.
		..#
		###`: "112",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(17, 1, in)
		if out != expectedOut {
			t.Errorf("day17 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay17Part2(t *testing.T) {

	testInputs := map[string]string{
		`.#.
		..#
		###`: "848",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(17, 2, in)
		if out != expectedOut {
			t.Errorf("day17 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
