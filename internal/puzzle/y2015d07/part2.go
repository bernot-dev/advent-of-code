package y2015d07

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
	components := make(map[string]component)
	for _, line := range in {
		c := NewComponent(line)
		components[c.output] = c
	}
	a1 := determineSignal(make(map[string]uint16), components, "a")
	a2 := determineSignal(map[string]uint16{"b": a1}, components, "a")
	return strconv.Itoa(int(a2)), nil
}
