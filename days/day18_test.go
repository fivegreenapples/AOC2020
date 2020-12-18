package days

import "testing"

func TestDay18Part1(t *testing.T) {

	testInputs := map[string]string{
		`1 + 2 * 3 + 4 * 5 + 6`:                           "71",
		`1 + (2 * 3) + (4 * (5 + 6))`:                     "51",
		`2 * 3 + (4 * 5)`:                                 "26",
		`5 + (8 * 3 + 9 + 3 * 4 * 3)`:                     "437",
		`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`:       "12240",
		`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`: "13632",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(18, 1, in)
		if out != expectedOut {
			t.Errorf("day18 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay18Part2(t *testing.T) {

	testInputs := map[string]string{
		`1 + 2 * 3 + 4 * 5 + 6`:                           "231",
		`1 + (2 * 3) + (4 * (5 + 6))`:                     "51",
		`2 * 3 + (4 * 5)`:                                 "46",
		`5 + (8 * 3 + 9 + 3 * 4 * 3)`:                     "1445",
		`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`:       "669060",
		`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`: "23340",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(18, 2, in)
		if out != expectedOut {
			t.Errorf("day18 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
