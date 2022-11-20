package y2015d02

import "testing"

func TestPart1(t *testing.T) {
	type test struct {
		input any
		want  any
	}

	tests := map[string]test{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := part1(tc.input)
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestPaperNeeded(t *testing.T) {
	type test struct {
		l, w, h int
		want    int
	}

	tests := map[string]test{
		"2x3x4": {
			l:    2,
			w:    3,
			h:    4,
			want: 58,
		},
		"1x1x10": {
			l:    1,
			w:    1,
			h:    10,
			want: 43,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := paperNeeded(tc.l, tc.w, tc.h)
			if got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
