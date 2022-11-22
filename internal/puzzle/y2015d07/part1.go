package y2015d07

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

type component struct {
	description    string
	componentType  string
	active         bool
	input          map[string]uint16
	output         string
	transformation func(...uint16) uint16
	value          uint16
}

func (c *component) connectInput(source string, val uint16) {
	c.input[source] = val
	var inputs []uint16
	switch c.componentType {
	case "AND", "OR":
		if len(c.input) < 2 {
			return
		}
		for _, v := range c.input {
			inputs = append(inputs, v)
		}
	default:
		inputs = append(inputs, val)
	}
	nVal := c.transformation(inputs...)
	if nVal != c.value {
		// connect input
	}
	c.value = nVal
}

func mustAtoi(a string) uint16 {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic("failed to covert to int: " + a)
	}
	return uint16(i)
}

func NewComponent(desc string) *component {
	c := component{
		description: desc,
		input:       make(map[string]uint16),
	}

	parts := strings.Split(desc, " -> ")
	componentTypeMatcher := regexp.MustCompile(`[A-Z]+`)
	c.componentType = componentTypeMatcher.FindString(parts[0])
	c.output = parts[1]

	numericMatcher := regexp.MustCompile(`\d+`)
	numeric := numericMatcher.FindString(parts[0])

	var n uint16
	if len(numeric) > 0 {
		n = mustAtoi(numeric)
	}

	switch c.componentType {
	case "":
		c.value = n
	case "AND":
		c.transformation = and
	case "OR":
		c.transformation = or
	case "LSHIFT":
		c.transformation = lshift(n)
	case "RSHIFT":
		c.transformation = rshift(n)
	case "NOT":
		c.transformation = not
	default:
		panic("unhandled component")
	}
	fmt.Println(c)
	return &c
}

func and(in ...uint16) uint16 {
	return in[0] & in[1]
}

func or(in ...uint16) uint16 {
	return in[0] | in[1]
}

func lshift(places uint16) func(...uint16) uint16 {
	return func(in ...uint16) uint16 {
		return in[0] << places
	}
}

func rshift(places uint16) func(...uint16) uint16 {
	return func(in ...uint16) uint16 {
		return in[0] >> places
	}
}

func not(in ...uint16) uint16 {
	return ^in[0]
}

func init() {
	register(1, Part1)
}

func Part1(r io.ReadCloser) (string, error) {
	in, err := input.AsStrings(r)
	if err != nil {
		return "", err
	}

	components := make(map[string]*component)
	for _, line := range in {
		c := NewComponent(line)
		components[line] = c
	}
	_ = components
	return "", nil
}
