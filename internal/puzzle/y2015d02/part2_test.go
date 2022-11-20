package y2015d02

import "testing"

func TestPart2(t *testing.T) {
	type test struct {
		input any
		want  any
	}

	tests := map[string]test{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := part2(tc.input)
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
