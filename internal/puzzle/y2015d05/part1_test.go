package y2015d05

import "testing"

func TestNice(t *testing.T) {
	tests := map[string]bool{
		"ugknbfddgicrmopn": true,
		"aaa":              true,
		"jchzalrnumimnmhp": false,
		"haegwjzuvuyypxyu": false,
		"dvszwmarrgswjxmb": false,
	}
	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := nice(tc, vowelCount(3), doubleLetter, bannedFree("ab", "cd", "pq", "xy"))
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}

func TestDoubleLetter(t *testing.T) {
	type test struct {
		count int
		want  bool
	}

	tests := map[string]test{
		"xx": {
			count: 2,
			want:  true,
		},
		"abcdde": {
			count: 2,
			want:  true,
		},
		"aabbccdd": {
			count: 2,
			want:  true,
		},
		"dvszwmarrgswjxmb": {
			count: 2,
			want:  true,
		},
	}

	for s, tc := range tests {
		t.Run(s, func(t *testing.T) {
			got := doubleLetter(s)
			if got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func TestVowelCount(t *testing.T) {
	type test struct {
		count int
		want  bool
	}

	tests := map[string]test{
		"aei": {
			count: 3,
			want:  true,
		},
		"xazegov": {
			count: 3,
			want:  true,
		},
		"aeiouaeiouaeiou": {
			count: 3,
			want:  true,
		},
		"abcde": {
			count: 3,
			want:  false,
		},
		"dvszwmarrgswjxmb": {
			count: 3,
			want:  false,
		},
	}

	for s, tc := range tests {
		t.Run(s, func(t *testing.T) {
			got := vowelCount(tc.count)(s)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestBannedFree(t *testing.T) {
	banned := []string{"ab", "cd", "pq", "xy"}
	tests := map[string]bool{
		"ugknbfddgicrmopn": true,
		"aaa":              true,
		"jchzalrnumimnmhp": true,
		"haegwjzuvuyypxyu": false,
		"dvszwmarrgswjxmb": true,
	}
	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := bannedFree(banned...)(tc)
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
