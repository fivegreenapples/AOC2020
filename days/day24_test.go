package days

import "testing"

func TestDay24Part1(t *testing.T) {

	testInputs := map[string]string{
		`esenee`:  "1",
		`esew`:    "1",
		`nwwswee`: "1",
		`sesenwnenenewseeswwswswwnenewsewsw
		neeenesenwnwwswnenewnwwsewnenwseswesw
		seswneswswsenwwnwse
		nwnwneseeswswnenewneswwnewseswneseene
		swweswneswnenwsewnwneneseenw
		eesenwseswswnenwswnwnwsewwnwsene
		sewnenenenesenwsewnenwwwse
		wenwwweseeeweswwwnwwe
		wsweesenenewnwwnwsenewsenwwsesesenwne
		neeswseenwwswnwswswnw
		nenwswwsewswnenenewsenwsenwnesesenew
		enewnwewneswsewnwswenweswnenwsenwsw
		sweneswneswneneenwnewenewwneswswnese
		swwesenesewenwneswnwwneseswwne
		enesenwswwswneneswsenwnewswseenwsese
		wnwnesenesenenwwnenwsewesewsesesew
		nenewswnwewswnenesenwnesewesw
		eneswnwswnwsenenwnwnwwseeswneewsenese
		neswnwewnwnwseenwseesewsenwsweewe
		wseweeenwnesenwwwswnew`: "10",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(24, 1, in)
		if out != expectedOut {
			t.Errorf("day24 pt1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay24Part2(t *testing.T) {

	testInputs := map[string]string{
		`sesenwnenenewseeswwswswwnenewsewsw
		neeenesenwnwwswnenewnwwsewnenwseswesw
		seswneswswsenwwnwse
		nwnwneseeswswnenewneswwnewseswneseene
		swweswneswnenwsewnwneneseenw
		eesenwseswswnenwswnwnwsewwnwsene
		sewnenenenesenwsewnenwwwse
		wenwwweseeeweswwwnwwe
		wsweesenenewnwwnwsenewsenwwsesesenwne
		neeswseenwwswnwswswnw
		nenwswwsewswnenenewsenwsenwnesesenew
		enewnwewneswsewnwswenweswnenwsenwsw
		sweneswneswneneenwnewenewwneswswnese
		swwesenesewenwneswnwwneseswwne
		enesenwswwswneneswsenwnewswseenwsese
		wnwnesenesenenwwnenwsewesewsesesew
		nenewswnwewswnenesenwnesewesw
		eneswnwswnwsenenwnwnwwseeswneewsenese
		neswnwewnwnwseenwseesewsenwsweewe
		wseweeenwnesenwwwswnew`: "2208",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(24, 2, in)
		if out != expectedOut {
			t.Errorf("day24 pt2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
