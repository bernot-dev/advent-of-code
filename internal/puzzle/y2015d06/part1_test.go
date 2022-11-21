package y2015d06

import "testing"

func TestParseInstruction(t *testing.T) {
	tests := map[string]instruction{
		"toggle 461,550 through 564,900": {
			op: "toggle",
			a:  coordinate{461, 550},
			b:  coordinate{564, 900},
		},
	}
	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := parseInstruction(tc)
			if got != want {
				t.Errorf("got: %+v\nwant: %+v", got, want)
			}
		})
	}
}

func TestApplyRange(t *testing.T) {
	tests := map[instruction]int{
		instruction{"turn on", coordinate{0, 0}, coordinate{999, 999}}: 1000000,
	}
	for tc, want := range tests {
		t.Run(tc.String(), func(t *testing.T) {
			g := make(grid)
			g.applyRange(tc.op, tc.a, tc.b)
			got := g.len()
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
