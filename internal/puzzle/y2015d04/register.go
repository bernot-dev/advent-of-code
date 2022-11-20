package y2015d04

import (
	"github.com/bernot-dev/advent-of-code/internal/puzzle"
	"github.com/bernot-dev/advent-of-code/internal/solver"
)

const (
	year = 2015
	day  = 4
)

func register(part int, fn solver.Solver) {
	ref := puzzle.Reference{
		Year: year,
		Day:  day,
		Part: part,
	}
	solver.Register(ref, fn)
}
