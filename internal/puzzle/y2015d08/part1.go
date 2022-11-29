package y2015d08

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

func Part1(r io.ReadCloser) (string, error) {
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}
	count := part1(in)
	return strconv.Itoa(count), nil
}

func part1(in []string) int {
	var count int
	for _, line := range in {
		count += len(line)
		unquoted, err := strconv.Unquote(line)
		if err != nil {
			panic("failed to unquote")
		}
		count -= len(unquoted)
	}
	return count
}
