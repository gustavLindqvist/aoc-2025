package day04

import (
	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star2(input string) any {
	var grid = make([][]bool, 0)
	for _, line := range aoc.Lines(input) {
		var g = make([]bool, 0)
		for _, r := range line {
			g = append(g, r == '@')
		}
		grid = append(grid, g)
	}
	var count = 0
	for true {
		var grid2 = make([][]bool, len(grid))
		var cd = 0
		for y := 0; y < len(grid); y++ {
			var g = make([]bool, len(grid[0]))
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] && (neighbours(grid, x, y) < 4) {
					cd += 1
					g[x] = false
				} else {
					g[x] = grid[y][x]
				}
			}
			grid2[y] = g
		}
		if cd == 0 {
			return count
		} else {
			count += cd
		}
		grid = grid2
	}

	return count
}
