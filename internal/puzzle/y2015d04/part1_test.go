package y2015d04

import (
	"encoding/hex"
	"testing"
)

func TestPart1(t *testing.T) {
	type test struct {
		input []byte
		want  int
	}

	tests := map[string]test{
		"abcdef": {
			input: []byte("abcdef"),
			want:  609043,
		},
		"pqrstuv": {
			input: []byte("pqrstuv"),
			want:  1048970,
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

func TestLeadingZeroes(t *testing.T) {
	type test struct {
		input  [16]byte
		zeroes int
		want   bool
	}

	example1Hash, err := hex.DecodeString("000001dbbfa3a5c83a2d506429c7b00e")
	if err != nil {
		t.Fatal(err)
	}
	example2Hash, err := hex.DecodeString("000006136ef2ff3b291c85725f17325c")
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]test{
		"0000012345678901": {
			input:  [16]byte{0x00, 0x00, 0x0F, 0xED, 0xCB, 0xA0, 0x98, 0x76},
			zeroes: 5,
			want:   true,
		},
		"ABCDEF0123456789": {
			input:  [16]byte{0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89},
			zeroes: 5,
			want:   false,
		},
		"0000ABCDEFG1234567890": {
			input:  [16]byte{0x00, 0x00, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45},
			zeroes: 5,
			want:   false,
		},
		"000001dbbfa3a5c83a2d506429c7b00e": {
			input:  *(*[16]byte)(example1Hash),
			zeroes: 5,
			want:   true,
		},
		"000006136ef2ff3b291c85725f17325c": {
			input:  *(*[16]byte)(example2Hash),
			zeroes: 5,
			want:   true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := leadingZeroes(tc.input, tc.zeroes)
			want := tc.want
			if got != want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
