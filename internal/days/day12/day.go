package day12

import (
	_ "embed"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

//go:embed input.txt
var Input string

//go:embed example.txt
var Example string

// Day returns the solvers for this day.
var Day = aoc.Day{
	Star1: Star1,
	Star2: Star2,
}
