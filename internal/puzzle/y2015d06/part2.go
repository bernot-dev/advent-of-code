package y2015d06

import (
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(2, Part2)
}

type dimmableGrid map[coordinate]int

func (g dimmableGrid) turnOn(c coordinate) int {
	g[c]++
	return 1
}

func (g dimmableGrid) turnOff(c coordinate) int {
	var diff int
	if g[c] > 0 {
		diff = -1
		g[c] -= 1
	}
	return diff
}

func (g dimmableGrid) toggle(c coordinate) int {
	g[c] += 2
	return 2
}

func (g dimmableGrid) applyRange(ins instruction) int {
	var op func(coordinate) int
	switch ins.op {
	case "turn on":
		op = g.turnOn
	case "turn off":
		op = g.turnOff
	case "toggle":
		op = g.toggle
	}
	var diff int
	for i := ins.a.x; i <= ins.b.x; i++ {
		for j := ins.a.y; j <= ins.b.y; j++ {
			diff += op(coordinate{i, j})
		}
	}
	return diff
}

func Part2(r io.ReadCloser) (string, error) {
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}
	instructions := make([]instruction, len(in))
	for i, line := range in {
		instructions[i] = parseInstruction(line)
	}
	return strconv.Itoa(part2(instructions)), nil
}

func part2(input []instruction) int {
	lights := make(dimmableGrid)
	var brightness int
	for _, i := range input {
		brightness += lights.applyRange(i)
	}
	return brightness
}
