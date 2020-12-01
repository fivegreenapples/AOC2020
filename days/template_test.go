package days

import "testing"

func TestDay___Part1(t *testing.T) {

	testInputs := map[string]string{}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 1, in)
		if out != expectedOut {
			t.Errorf("day___ pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay___Part2(t *testing.T) {

	testInputs := map[string]string{}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 2, in)
		if out != expectedOut {
			t.Errorf("day___ pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
