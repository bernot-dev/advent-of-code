package y2015d05

import (
	"io"
	"strconv"
	"strings"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

type Validator func(string) bool

func Part1(r io.ReadCloser) (string, error) {
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(part1(in)), nil
}

func part1(input []string) int {
	var niceCount int
	for _, s := range input {
		if nice(s, vowelCount(3), doubleLetter, bannedFree("ab", "cd", "pq", "xy")) {
			niceCount++
		}
	}
	return niceCount
}

func nice(s string, validators ...Validator) bool {
	for _, valid := range validators {
		if !valid(s) {
			return false
		}
	}
	return true
}

func doubleLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}

func vowelCount(n int) func(string) bool {
	var vowels int
	return func(s string) bool {
		for _, r := range s {
			switch r {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			default:
				continue
			}
			if vowels >= n {
				return true
			}
		}
		return false
	}
}

func bannedFree(banned ...string) func(string) bool {
	return func(s string) bool {
		for _, b := range banned {
			if strings.Contains(s, b) {
				return false
			}
		}
		return true
	}
}
