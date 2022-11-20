package y2015d04

import (
	"crypto/md5"
	"io"
	"strconv"

	"github.com/bernot-dev/advent-of-code/internal/input"
)

func init() {
	register(2, Part2)
}

func Part2(r io.ReadCloser) (string, error) {
	b, err := input.AsBytes(r)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(part2(b)), nil
}

func part2(input []byte) int {
	for n := 1; true; n++ {
		numBytes := []byte(strconv.Itoa(n))
		b := append(input, numBytes...)
		h := md5.Sum(b)
		if leadingZeroes(h, 6) {
			return n
		}
	}
	panic("no hash found")
}
