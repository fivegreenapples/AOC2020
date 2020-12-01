package days

import "testing"

func TestDay1Part1(t *testing.T) {

	testInputs := map[string]string{
		`1721
		979
		366
		299
		675
		1456`: "514579",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 1, in)
		if out != expectedOut {
			t.Errorf("day1 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay1Part2(t *testing.T) {

	testInputs := map[string]string{
		`1721
		979
		366
		299
		675
		1456`: "241861950",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 2, in)
		if out != expectedOut {
			t.Errorf("day1 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
