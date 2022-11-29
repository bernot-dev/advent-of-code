package y2015d07

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

type gate func(...uint16) uint16

func Part1(r io.ReadCloser) (string, error) {
	memo := make(map[string]uint16)
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}
	m := make(map[string]component)
	for _, line := range in {
		c := NewComponent(line)
		m[c.output] = c
	}
	return strconv.Itoa(int(determineSignal(memo, m, "a"))), nil
}

func determineSignal(memo map[string]uint16, components map[string]component, wire string) uint16 {
	if n, ok := memo[wire]; ok {
		return n
	}
	if n, err := strconv.Atoi(wire); err == nil {
		memo[wire] = uint16(n)
		return memo[wire]
	}
	c := components[wire]
	signals := make([]uint16, 0)
	for _, wire := range c.inputs {
		signals = append(signals, determineSignal(memo, components, wire))
	}
	memo[wire] = c.op(signals...)
	return memo[wire]
}
