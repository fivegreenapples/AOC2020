package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/fivegreenapples/AOC2020/days"
)

func main() {

	day := flag.Int("d", 0, "Day of Advent")
	verbose := flag.Bool("v", false, "Verbosity")
	part := flag.Int("p", 0, "Part")
	cpuprofile := flag.String("profile", "", "File for cpu profile")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Printf("Error: couldn't create cpu profile: %v\n", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *day <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	var puzzleInput string

	inputFile := "inputs/day" + strconv.Itoa(*day) + ".txt"
	puzzleInputBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error: couldn't read input file: %v\n", err)
		os.Exit(2)
	}
	puzzleInput = string(puzzleInputBytes)

	runner := days.NewRunner(*verbose)

	if *part == 0 || *part == 1 {
		pt1Out, pt1err := runner.Run(*day, 1, puzzleInput)
		if pt1err != nil {
			fmt.Printf("%s\n", pt1err.Error())
			os.Exit(3)
		}
		fmt.Println(pt1Out)
	}
	if *part == 0 || *part == 2 {
		pt2Out, pt2err := runner.Run(*day, 2, puzzleInput)
		if pt2err == nil {
			fmt.Println(pt2Out)
		} else if *part == 2 {
			// only display error if we asked for part 2
			fmt.Printf("%s\n", pt2err.Error())
			os.Exit(3)
		}
	}
}
