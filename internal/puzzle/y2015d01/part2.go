package y2015d01

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(2, Part2)
}

func Part2(r io.ReadCloser) (string, error) {
	in, err := input.AsString(r)
	if err != nil {
		return "", err
	}
	solution := part2(in)
	return strconv.Itoa(solution), nil
}

func part2(input string) int {
	var floor int
	for i, r := range input {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return floor
}
