package days

import "testing"

func TestDay11Part1(t *testing.T) {

	testInputs := map[string]string{
		`L.LL.LL.LL
		LLLLLLL.LL
		L.L.L..L..
		LLLL.LL.LL
		L.LL.LL.LL
		L.LLLLL.LL
		..L.L.....
		LLLLLLLLLL
		L.LLLLLL.L
		L.LLLLL.LL`: "37",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(11, 1, in)
		if out != expectedOut {
			t.Errorf("day11 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay11Part2(t *testing.T) {

	testInputs := map[string]string{
		`L.LL.LL.LL
		LLLLLLL.LL
		L.L.L..L..
		LLLL.LL.LL
		L.LL.LL.LL
		L.LLLLL.LL
		..L.L.....
		LLLLLLLLLL
		L.LLLLLL.L
		L.LLLLL.LL`: "26",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(11, 2, in)
		if out != expectedOut {
			t.Errorf("day11 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
