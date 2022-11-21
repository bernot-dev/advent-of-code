package y2015d05

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(2, Part2)
}

func Part2(r io.ReadCloser) (string, error) {
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(part2(in)), nil
}

func part2(input []string) int {
	var niceCount int
	for _, s := range input {
		if repeatedPair(s) && letterSandwich(s) {
			niceCount++
		}
	}
	return niceCount
}

func repeatedPair(s string) bool {
	firstOccurrence := make(map[string]int)
	for i := 0; i <= len(s)-2; i++ {
		currentSequence := s[i : i+2]
		if _, ok := firstOccurrence[currentSequence]; !ok {
			firstOccurrence[currentSequence] = i
			continue
		}
		if i >= firstOccurrence[currentSequence]+2 {
			return true
		}
	}
	return false
}

func letterSandwich(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}
