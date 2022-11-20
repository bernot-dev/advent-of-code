package y2015d02

import (
	"github.com/bernot-dev/advent-of-code/internal/puzzle"
	"github.com/bernot-dev/advent-of-code/internal/solver"
)

const (
	year = 2015
	day  = 2
)

func register(part int, fn solver.Solver) {
	ref := puzzle.Reference{
		Year: year,
		Day:  day,
		Part: part,
	}
	solver.Register(ref, fn)
}
