package days

import "testing"

func TestDay6Part1(t *testing.T) {

	testInputs := map[string]string{
		`abcx
abcy
abcz`: "6",
		`abc

a
b
c

ab
ac

a
a
a
a

b`: "11",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(6, 1, in)
		if out != expectedOut {
			t.Errorf("day6 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay6Part2(t *testing.T) {

	testInputs := map[string]string{
		`abc

a
b
c

ab
ac

a
a
a
a

b`: "6",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(6, 2, in)
		if out != expectedOut {
			t.Errorf("day6 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
