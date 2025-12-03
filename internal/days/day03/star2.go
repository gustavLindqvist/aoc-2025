package day03

import (
	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star2(input string) any {
	var count = 0
	for _, bank := range aoc.Lines(input) {
		var bats = make([]int, 12)
		var m = len(bank) - 1
		for i, r := range bank {
			var jolt = int(r - '0')
			var start = max(0, 11-(m-i))
			for j := start; j < 12; j++ {
				if jolt > bats[j] {
					bats[j] = jolt
					clear(bats[(j + 1):])
					break
				}
			}
		}
		var pow = 1
		for i := 11; i >= 0; i-- {
			count += bats[i] * pow
			pow *= 10
		}
	}
	return count
}
