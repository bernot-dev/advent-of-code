package y2015d01

import (
	"github.com/bernot-dev/advent-of-code/internal/puzzle"
	"github.com/bernot-dev/advent-of-code/internal/solver"
)

const (
	year = 2015
	day  = 1
)

func register(part int, fn solver.Solver) {
	ref := puzzle.Reference{
		Year: year,
		Day:  day,
		Part: part,
	}
	solver.Register(ref, fn)
}
