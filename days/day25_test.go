package days

import "testing"

func TestDay25Part1(t *testing.T) {

	testInputs := map[string]string{
		"5764801\n17807724": "14897079",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(25, 1, in)
		if out != expectedOut {
			t.Errorf("day25 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay25Part2(t *testing.T) {

	testInputs := map[string]string{
		``: "0",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(25, 2, in)
		if out != expectedOut {
			t.Errorf("day25 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
