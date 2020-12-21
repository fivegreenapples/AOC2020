package days

import "testing"

func TestDay21Part1(t *testing.T) {

	testInputs := map[string]string{
		`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
		trh fvjkl sbzzf mxmxvkd (contains dairy)
		sqjhc fvjkl (contains soy)
		sqjhc mxmxvkd sbzzf (contains fish)`: "5",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(21, 1, in)
		if out != expectedOut {
			t.Errorf("day21 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay21Part2(t *testing.T) {

	testInputs := map[string]string{
		`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
		trh fvjkl sbzzf mxmxvkd (contains dairy)
		sqjhc fvjkl (contains soy)
		sqjhc mxmxvkd sbzzf (contains fish)`: "mxmxvkd,sqjhc,fvjkl",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(21, 2, in)
		if out != expectedOut {
			t.Errorf("day21 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
