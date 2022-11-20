package y2015d03

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
	m := make(map[Location]int)
	s := Location{}
	r := Location{}
	m[s]++
	m[r]++
	for i, v := range input {
		var l *Location
		switch i%2 == 0 {
		case true:
			l = &s
		case false:
			l = &r
		}
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
		m[*l]++
	}
	return len(m)
}
