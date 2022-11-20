package y2015d04

import (
	"crypto/md5"
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(1, Part1)
}

func Part1(r io.ReadCloser) (string, error) {
	b, err := input.AsBytes(r)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(part1(b)), nil
}

func part1(input []byte) int {
	for n := 1; true; n++ {
		numBytes := []byte(strconv.Itoa(n))
		b := append(input, numBytes...)
		h := md5.Sum(b)
		if leadingZeroes(h, 5) {
			return n
		}
	}
	panic("no hash found")
}

func leadingZeroes(b [16]byte, n int) bool {
	for i := 0; i < n/2; i++ {
		if b[i] != 0 {
			return false
		}
	}
	if n%2 == 1 {
		return b[n/2]&0xF0 == 0
	}
	return true
}
