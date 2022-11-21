package y2015d05

import "testing"

func TestPart2(t *testing.T) {
	type test struct {
		input []string
		want  int
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

func TestNiceV2(t *testing.T) {
	tests := map[string]bool{
		"qjhvhtzxzqqjkmpb": true,
		"xxyxx":            true,
		"uurcxstgmygtbstg": false,
		"ieodomkazucvgmuy": false,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := nice(tc, repeatedPair, letterSandwich)
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}

func TestRepeatedPair(t *testing.T) {
	tests := map[string]bool{
		"xyxy":             true,
		"aabcdefgaa":       true,
		"aaa":              false,
		"qjhvhtzxzqqjkmpb": true,
		"xxyxx":            true,
		"uurcxstgmygtbstg": true,
		"ieodomkazucvgmuy": false,
		"bbcbb":            true,
		"rxexcbwhiywwwwnu": true,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := repeatedPair(tc)
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}

func TestLetterSandwich(t *testing.T) {
	tests := map[string]bool{
		"xyx":              true,
		"abcdefeghi":       true,
		"aaa":              true,
		"qjhvhtzxzqqjkmpb": true,
		"xxyxx":            true,
		"uurcxstgmygtbstg": false,
		"ieodomkazucvgmuy": true,
		"rxexcbwhiywwwwnu": true,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := letterSandwich(tc)
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
