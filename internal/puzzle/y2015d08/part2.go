package y2015d08

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
	count := part2(in)
	return strconv.Itoa(count), nil
}

func part2(in []string) int {
	var count int
	for _, line := range in {
		quoted := strconv.Quote(line)
		count += len(quoted)
		count -= len(line)
	}
	return count
}
