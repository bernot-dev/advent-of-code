package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bernot-dev/advent-of-code/internal/aoc"
	"github.com/bernot-dev/advent-of-code/internal/puzzle"
	"github.com/bernot-dev/advent-of-code/internal/solver"
	flag "github.com/spf13/pflag"
)

func main() {
	// Determine which challenge to run
	y, _, d := time.Now().Date()
	var year *int = flag.IntP("year", "y", y, "Year for Advent of Code challenge")
	var day *int = flag.IntP("day", "d", d, "Day for Advent of Code Challenge")
	flag.Parse()

	log.Printf("Running: %d-%02d\n", *year, *day)

	for _, part := range []int{1, 2} {
		in, err := input(*year, *day)
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()

		ref := puzzle.Reference{
			Year: *year,
			Day:  *day,
			Part: part,
		}
		solution, err := solver.Solve(ref, in)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v Solution: %s\n", ref, solution)
	}
}

func input(year, day int) (io.ReadCloser, error) {
	var inputFileExists bool
	if _, err := os.Stat(inputFilename(year, day)); err == nil {
		log.Printf("Input file exists locally: %q\n", inputFilename(year, day))
		inputFileExists = true
	}

	if !inputFileExists {
		log.Printf("Input file does not exist locally, getting from AOC\n")

		// Get input from web
		aocReader, err := aoc.Input(year, day)
		if err != nil {
			return nil, err
		}
		input, err := io.ReadAll(aocReader)
		if err != nil {
			return nil, err
		}

		// Save to local file
		os.WriteFile(inputFilename(year, day), input, 0400)
	}

	// Use local input file, if available
	f, err := inputFile(year, day)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func inputFile(year, day int) (io.ReadCloser, error) {
	f, err := os.Open(inputFilename(year, day))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func puzzleRef(year, day, part int) string {
	return fmt.Sprintf("%d%02d%d", year, day, part)
}

func inputFilename(year, day int) string {
	return fmt.Sprintf("input/%d-%02d.txt", year, day)
}
