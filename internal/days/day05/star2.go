package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star2(input string) any {
	var parts = strings.Split(input, "\n\n")
	var first, _ = parts[0], parts[1]
	var inv = make([]fresh, 0)

	for _, line := range aoc.Lines(first) {
		var tmp = strings.Split(line, "-")
		var first, _ = strconv.Atoi(tmp[0])
		var second, _ = strconv.Atoi(tmp[1])
		inv = append(inv, fresh{first, second})
	}

	sort.Slice(inv, func(i, j int) bool {
		if inv[i].from != inv[j].from {
			return inv[i].from < inv[j].from
		}
		return inv[i].to < inv[j].to
	})
	var count = 0
	var lastEnd = 0
	for _, f := range inv {
		if f.from > lastEnd {
			count += 1
		}
		var from = max(lastEnd, f.from)
		var to = max(lastEnd, f.to)
		count += to - from
		lastEnd = to
	}
	return count
}
