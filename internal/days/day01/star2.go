package day01

import (
	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
	"strconv"
)

func Star2(input string) any {
	pos := 50
	count := 0

	for _, line := range aoc.Lines(input) {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])
		lo, hi := 0, 0

		switch dir {
		case 'L':
			lo = pos - steps
			hi = pos - 1
			pos = ((pos-steps)%100 + 100) % 100
		case 'R':
			lo = pos + 1
			hi = pos + steps
			pos = ((pos+steps)%100 + 100) % 100
		}

		if lo <= hi {
			count += divDown(hi) - divDown(lo-1)
		}
	}

	return count
}

func divDown(a int) int {
	if a >= 0 {
		return a / 100
	}
	return (a - 99) / 100
}
