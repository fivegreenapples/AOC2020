package days

import "testing"

func TestDay15Part1(t *testing.T) {

	testInputs := map[string]string{
		`0,3,6`: "436",
		`1,3,2`: "1",
		`2,1,3`: "10",
		`1,2,3`: "27",
		`2,3,1`: "78",
		`3,2,1`: "438",
		`3,1,2`: "1836",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(15, 1, in)
		if out != expectedOut {
			t.Errorf("day15 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay15Part2(t *testing.T) {

	testInputs := map[string]string{
		`0,3,6`: "175594",
		`1,3,2`: "2578",
		`2,1,3`: "3544142",
		`1,2,3`: "261214",
		`2,3,1`: "6895259",
		`3,2,1`: "18",
		`3,1,2`: "362",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(15, 2, in)
		if out != expectedOut {
			t.Errorf("day15 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
