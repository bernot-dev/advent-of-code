package solver

import (
	"fmt"
	"io"
	"sync"

	"github.com/bernot-dev/advent-of-code/internal/puzzle"
)

type Solver func(io.ReadCloser) (string, error)

var (
	solversMutex sync.RWMutex
	solvers      = make(map[puzzle.Reference]Solver)
)

func Register(ref puzzle.Reference, fn Solver) {
	solversMutex.Lock()
	defer solversMutex.Unlock()
	solvers[ref] = fn
}

func Solve(ref puzzle.Reference, r io.ReadCloser) (string, error) {
	solversMutex.RLock()
	defer solversMutex.RUnlock()
	if _, ok := solvers[ref]; !ok {
		return "", fmt.Errorf("no solver registered for %+v", ref)
	}
	return solvers[ref](r)
}
