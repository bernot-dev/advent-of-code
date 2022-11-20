package y2015d03

import "testing"

func TestPart1(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := map[string]test{
		">": {
			input: ">",
			want:  2,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := part1(tc.input)
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
