package days

import "testing"

func TestDay5Part1(t *testing.T) {

	testInputs := map[string]string{
		`FBFBBFFRLR`: "357",
		`BFFFBBFRRR`: "567",
		`FFFBBBFRRR`: "119",
		`BBFFBBFRLL`: "820",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(5, 1, in)
		if out != expectedOut {
			t.Errorf("day5 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
