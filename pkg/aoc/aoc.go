// Package aoc provides common utilities for Advent of Code solutions.
package aoc

import (
	"strings"
)

// Lines splits input into lines, trimming trailing whitespace.
func Lines(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func LineBytes(input string) [][]byte {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	out := make([][]byte, len(lines))
	for i, line := range lines {
		out[i] = []byte(line)
	}
	return out
}

// Solver is the interface that each star solution must implement.
type Solver func(input string) any

// Day represents a single day's solutions.
type Day struct {
	Star1 Solver
	Star2 Solver
}
