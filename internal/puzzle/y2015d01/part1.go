package y2015d01

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

func Part1(r io.ReadCloser) (string, error) {
	in, err := input.AsString(r)
	if err != nil {
		return "", err
	}
	solution := part1(in)
	return strconv.Itoa(solution), nil
}

func part1(input string) int {
	var floor int
	for _, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}
