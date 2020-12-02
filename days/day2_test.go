package days

import "testing"

func TestDay2Part1(t *testing.T) {

	testInputs := map[string]string{
		`1-3 a: abcde
		1-3 b: cdefg
		2-9 c: ccccccccc`: "2",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(2, 1, in)
		if out != expectedOut {
			t.Errorf("day2 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay2Part2(t *testing.T) {

	testInputs := map[string]string{
		`1-3 a: abcde
		1-3 b: cdefg
		2-9 c: ccccccccc`: "1",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(2, 2, in)
		if out != expectedOut {
			t.Errorf("day2 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
