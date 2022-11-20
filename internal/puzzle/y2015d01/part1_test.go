package y2015d01

import "testing"

func TestPart1(t *testing.T) {
	tests := map[string]int{
		"(())":    0,
		"()()":    0,
		"(((":     3,
		"(()(()(": 3,
		"))(((((": 3,
		"())":     -1,
		"))(":     -1,
		")))":     -3,
		")())())": -3,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := part1(tc)
			if got != want {
				t.Errorf("case %q, want %d, got %d", tc, want, got)
			}
		})
	}
}
