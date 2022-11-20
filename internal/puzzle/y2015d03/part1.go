package y2015d03

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

const (
	N = '^'
	S = 'v'
	E = '>'
	W = '<'
)

type Location struct {
	Lat  int
	Long int
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
	m := make(map[Location]int)
	l := Location{}
	m[l]++
	for _, v := range input {
		switch v {
		case N:
			l.Lat++
		case S:
			l.Lat--
		case E:
			l.Long++
		case W:
			l.Long--
		}
		m[l]++
	}
	return len(m)
}
