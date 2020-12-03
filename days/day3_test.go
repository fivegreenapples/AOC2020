package days

import "testing"

func TestDay3Part1(t *testing.T) {

	testInputs := map[string]string{
		`..##.......
		#...#...#..
		.#....#..#.
		..#.#...#.#
		.#...##..#.
		..#.##.....
		.#.#.#....#
		.#........#
		#.##...#...
		#...##....#
		.#..#...#.#`: "7",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 1, in)
		if out != expectedOut {
			t.Errorf("day3 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay3Part2(t *testing.T) {

	testInputs := map[string]string{
		`..##.......
		#...#...#..
		.#....#..#.
		..#.#...#.#
		.#...##..#.
		..#.##.....
		.#.#.#....#
		.#........#
		#.##...#...
		#...##....#
		.#..#...#.#`: "336",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 2, in)
		if out != expectedOut {
			t.Errorf("day3 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
