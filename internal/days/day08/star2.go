package day08

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

func Star2(input string) any {
	var boxes = [][]int{}
	for _, l := range aoc.Lines(input) {
		var b = []int{}
		for i := range strings.SplitSeq(l, ",") {
			var num, _ = strconv.Atoi(i)
			b = append(b, num)
		}
		boxes = append(boxes, b)
	}

	dists := []pair{}
	for i, a := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			b := boxes[j]
			d := dist(a, b)
			dists = append(dists, pair{d, i, j})
		}
	}
	sort.Slice(dists, func(i, j int) bool {
		return dists[i].dist < dists[j].dist
	})
	clusters := make([]int, len(boxes))
	for i := range len(clusters) {
		clusters[i] = i
	}
	unique := len(clusters)
	for j := 0; j < len(dists); j++ {
		p := dists[j]
		a := clusters[p.a]
		b := clusters[p.b]
		if a != b {
			for i, val := range clusters {
				if val == b {
					clusters[i] = a
				}
			}
			unique -= 1
		}
		if unique == 1 {
			return boxes[p.a][0] * boxes[p.b][0]
		}
	}
	return nil
}
