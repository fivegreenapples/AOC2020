package days

import "testing"

func TestDay8Part1(t *testing.T) {

	testInputs := map[string]string{
		`nop +0
		acc +1
		jmp +4
		acc +3
		jmp -3
		acc -99
		acc +1
		jmp -4
		acc +6`: "5",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(8, 1, in)
		if out != expectedOut {
			t.Errorf("day8 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay8Part2(t *testing.T) {

	testInputs := map[string]string{
		`nop +0
		acc +1
		jmp +4
		acc +3
		jmp -3
		acc -99
		acc +1
		jmp -4
		acc +6`: "8",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(8, 2, in)
		if out != expectedOut {
			t.Errorf("day8 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
