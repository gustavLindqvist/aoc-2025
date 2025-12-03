package day03

import (
	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star1(input string) any {
	var count = 0
	for _, line := range aoc.Lines(input) {
		var f, s = 0, 0
		for i, r := range line {
			var b = int(r - '0')
			if b > f && i < len(line)-1 {
				f = b
				s = 0
			} else if b > s {
				s = b
			}
		}
		count += (f * 10) + s
	}
	return count
}
