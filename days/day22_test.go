package days

import "testing"

func TestDay22Part1(t *testing.T) {

	testInputs := map[string]string{
		`Player 1:
		9
		2
		6
		3
		1
		
		Player 2:
		5
		8
		4
		7
		10`: "306",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(22, 1, in)
		if out != expectedOut {
			t.Errorf("day22 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay22Part2(t *testing.T) {

	testInputs := map[string]string{
		`Player 1:
		9
		2
		6
		3
		1
		
		Player 2:
		5
		8
		4
		7
		10`: "291",
		`Player 1:
		43
		19
		
		Player 2:
		2
		29
		14`: "105",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(22, 2, in)
		if out != expectedOut {
			t.Errorf("day22 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
