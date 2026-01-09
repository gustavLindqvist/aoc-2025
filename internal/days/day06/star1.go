package day06

import (
	"strconv"
	"strings"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star1(input string) any {
	var lines = aoc.Lines(input)
	var ops = strings.Fields(lines[len(lines)-1])
	var nums = make([]int, len(ops))
	for i, field := range strings.Fields(lines[0]) {
		nums[i], _ = strconv.Atoi(field)
	}
	for _, l := range lines[1:(len(lines) - 1)] {
		for i, field := range strings.Fields(l) {
			var num, _ = strconv.Atoi(field)
			switch ops[i] {
			case "+":
				nums[i] += num
			case "*":
				nums[i] *= num
			}
		}
	}
	var result = 0
	for _, i := range nums {
		result += i
	}
	return result
}
