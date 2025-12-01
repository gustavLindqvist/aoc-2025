package day01

import (
	"strconv"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star1(input string) any {
	pos := 50
	count := 0

	for _, line := range aoc.Lines(input) {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		switch dir {
		case 'L':
			pos = ((pos-steps)%100 + 100) % 100
		case 'R':
			pos = ((pos+steps)%100 + 100) % 100
		}
		if pos == 0 {
			count += 1
		}
	}

	return count
}
