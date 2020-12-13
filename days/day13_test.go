package days

import "testing"

func TestDay13Part1(t *testing.T) {

	testInputs := map[string]string{
		`939
		7,13,x,x,59,x,31,19`: "295",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(13, 1, in)
		if out != expectedOut {
			t.Errorf("day13 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay13Part2(t *testing.T) {

	testInputs := map[string]string{
		`7,13,x,x,59,x,31,19`: "1068781",
		`17,x,13,19`:          "3417",
		`67,7,59,61`:          "754018",
		`67,x,7,59,61`:        "779210",
		`67,7,x,59,61`:        "1261476",
		`1789,37,47,1889`:     "1202161486",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(13, 2, "_\n"+in)
		if out != expectedOut {
			t.Errorf("day13 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
