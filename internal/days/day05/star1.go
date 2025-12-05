package day05

import (
	"strconv"
	"strings"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

type fresh struct {
	from int
	to   int
}

func Star1(input string) any {
	var parts = strings.Split(input, "\n\n")
	var first, second = parts[0], parts[1]
	var inv = make([]fresh, 0)
	for _, line := range aoc.Lines(first) {
		var tmp = strings.Split(line, "-")
		var first, _ = strconv.Atoi(tmp[0])
		var second, _ = strconv.Atoi(tmp[1])
		inv = append(inv, fresh{first, second})
	}
	var count = 0
	for _, line := range aoc.Lines(second) {
		var item, _ = strconv.Atoi(line)
		for _, f := range inv {
			if item >= f.from && item <= f.to {
				count += 1
				break
			}
		}
	}
	return count
}
