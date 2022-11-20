package y2015d02

import (
	"io"
	"strconv"
	"strings"

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

	var sum int
	for _, s := range in {
		parts := strings.Split(s, "x")
		var ints [3]int
		for j, v := range parts {
			ints[j], _ = strconv.Atoi(v)
		}
		sum += ribbonNeeded(ints[0], ints[1], ints[2])
	}
	return strconv.Itoa(sum), nil
}

func ribbonNeeded(l, w, h int) int {
	a, b := lowTwo(l, w, h)
	ribbon := 2 * (a + b)
	bow := l * w * h
	return ribbon + bow
}

func lowTwo(a, b, c int) (int, int) {
	if a >= b && a >= c {
		return b, c
	}
	if b >= a && b >= c {
		return a, c
	}
	if c >= a && c >= b {
		return a, b
	}
	panic("one side should be longer than the others")
}

func part2(input any) any {
	return nil
}
