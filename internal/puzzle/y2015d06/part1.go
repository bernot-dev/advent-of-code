package y2015d06

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

type instruction struct {
	op   string
	a, b coordinate
}

func (i *instruction) String() string {
	return fmt.Sprintf("%s %d,%d through %d,%d", i.op, i.a.x, i.a.y, i.b.x, i.b.y)
}

type coordinate struct {
	x, y int
}

type grid map[coordinate]bool

func (g grid) toggle(c coordinate) {
	if g[c] {
		delete(g, c)
		return
	}
	g[c] = true
}

func (g grid) applyRange(op string, a, b coordinate) {
	var o func(coordinate)
	switch op {
	case "toggle":
		o = g.toggle
	case "turn on":
		o = g.turnOn
	case "turn off":
		o = g.turnOff
	}
	for i := a.x; i <= b.x; i++ {
		for j := a.y; j <= b.y; j++ {
			o(coordinate{i, j})
		}
	}
}

func (g grid) turnOn(c coordinate) {
	g[c] = true
}

func (g grid) turnOff(c coordinate) {
	delete(g, c)
}

func (g grid) len() int {
	return len(g)
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert to int: " + s)
	}
	return n
}

func parseInstruction(s string) instruction {
	instructionRegexp := regexp.MustCompile(`(?P<op>toggle|turn off|turn on) (?P<a>\d+),(?P<b>\d+) through (?P<x>\d+),(?P<y>\d+)`)
	match := instructionRegexp.FindStringSubmatch(s)
	return instruction{
		op: match[1],
		a:  coordinate{mustAtoi(match[2]), mustAtoi(match[3])},
		b:  coordinate{mustAtoi(match[4]), mustAtoi(match[5])},
	}
}

func Part1(r io.ReadCloser) (string, error) {
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}
	instructions := make([]instruction, len(in))
	for i, line := range in {
		instructions[i] = parseInstruction(line)
	}
	return strconv.Itoa(part1(instructions)), nil
}

func part1(input []instruction) int {
	lights := make(grid)
	for _, i := range input {
		lights.applyRange(i.op, i.a, i.b)
	}
	return lights.len()
}
