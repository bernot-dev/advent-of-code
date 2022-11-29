package y2015d07

import (
	"strconv"
	"strings"
)

type component struct {
	inputs []string
	op     gate
	output string
}

var op = map[string]gate{
	"": func(n ...uint16) uint16 {
		return n[0]
	},
	"AND": func(n ...uint16) uint16 {
		return n[0] & n[1]
	},
	"OR": func(n ...uint16) uint16 {
		return n[0] | n[1]
	},
	"LSHIFT": func(n ...uint16) uint16 {
		return n[0] << n[1]
	},
	"RSHIFT": func(n ...uint16) uint16 {
		return n[0] >> n[1]
	},
	"NOT": func(n ...uint16) uint16 {
		return ^n[0]
	},
}

func NewComponent(line string) component {
	c := component{}
	parts := strings.Split(line, " -> ")
	gate, output := parts[0], parts[1]
	c.output = output
	gateParts := strings.Split(gate, " ")
	switch len(gateParts) {
	case 1:
		c.inputs = []string{gateParts[0]}
		c.op = op[""]
	case 2:
		c.inputs = []string{gateParts[1]}
		c.op = op[gateParts[0]]
	case 3:
		c.inputs = []string{gateParts[0], gateParts[2]}
		c.op = op[gateParts[1]]
	default:
		panic("unexpected length of gate parts")
	}
	return c
}

func mustAtoi(s string) uint16 {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint16(n)
}
