package day04

import "github.com/gustavLindqvist/aoc-2025/pkg/aoc"

func Star1(input string) any {
	var grid = make([][]bool, 0)
	for _, line := range aoc.Lines(input) {
		var g = make([]bool, 0)
		for _, r := range line {
			g = append(g, r == '@')
		}
		grid = append(grid, g)
	}
	var count = 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] && (neighbours(grid, x, y) < 4) {
				count += 1
			}
		}
	}

	return count
}

func neighbours(grid [][]bool, x int, y int) int {
	var count = 0
	for yy := max(0, y-1); yy < min(len(grid), y+2); yy++ {
		for xx := max(0, x-1); xx < min(len(grid[0]), x+2); xx++ {
			if grid[yy][xx] {
				count += 1
			}
		}
	}
	return (count - 1)
}
