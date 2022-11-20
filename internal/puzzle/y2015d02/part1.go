package y2015d02

import (
	"io"
	"strconv"
	"strings"

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

	var sum int
	for _, s := range in {
		parts := strings.Split(s, "x")
		var ints [3]int
		for j, v := range parts {
			ints[j], _ = strconv.Atoi(v)
		}
		sum += paperNeeded(ints[0], ints[1], ints[2])
	}
	return strconv.Itoa(sum), nil
}

func part1(input any) any {
	return nil
}

func paperNeeded(l, w, h int) int {
	lw, lh, wh := l*w, l*h, w*h
	s := min(lw, lh, wh)
	return 2*lw + 2*lh + 2*wh + s
}

func min(a ...int) int {
	var min = a[0]
	for _, n := range a {
		if n < min {
			min = n
		}
	}
	return min
}
