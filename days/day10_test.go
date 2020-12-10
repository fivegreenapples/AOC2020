package days

import "testing"

func TestDay10Part1(t *testing.T) {

	testInputs := map[string]string{
		`16
		10
		15
		5
		1
		11
		7
		19
		6
		12
		4`: "35",
		`28
		33
		18
		42
		31
		14
		46
		20
		48
		47
		24
		23
		49
		45
		19
		38
		39
		11
		1
		32
		25
		35
		8
		17
		7
		9
		4
		2
		34
		10
		3`: "220",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(10, 1, in)
		if out != expectedOut {
			t.Errorf("day10 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay10Part2(t *testing.T) {

	testInputs := map[string]string{
		`16
		10
		15
		5
		1
		11
		7
		19
		6
		12
		4`: "8",
		`28
		33
		18
		42
		31
		14
		46
		20
		48
		47
		24
		23
		49
		45
		19
		38
		39
		11
		1
		32
		25
		35
		8
		17
		7
		9
		4
		2
		34
		10
		3`: "19208",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(10, 2, in)
		if out != expectedOut {
			t.Errorf("day10 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
