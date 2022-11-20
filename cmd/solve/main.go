package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"

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
		log.Printf("Input file does not exist locally, getting from AOC: %q\n", inputURL(year, day))

		// Get input from web
		aocReader, err := inputAOC(year, day)
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

func inputAOC(year, day int) (io.ReadCloser, error) {
	session, err := os.ReadFile("aoc-session")
	if err != nil {
		return nil, err
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: strings.TrimSpace(string(session)),

		Path:    "/",
		Domain:  "adventofcode.com",
		Expires: time.Now().Add(time.Hour),

		MaxAge:   0,
		Secure:   true,
		HttpOnly: false,
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest(http.MethodGet, inputURL(year, day), nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&sessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		errorText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorText))
	}
	return resp.Body, nil
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

func inputURL(year, day int) string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
}
