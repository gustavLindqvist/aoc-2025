package day08

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

type pair struct {
	dist int
	a    int
	b    int
}

func Star1(input string) any {
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
	iterations := 1000
	if len(dists) < 1000 {
		iterations = 10
	}
	for j := 0; j < iterations; j++ {
		p := dists[j]
		a := clusters[p.a]
		b := clusters[p.b]
		if a != b {
			for i, val := range clusters {
				if val == b {
					clusters[i] = a
				}
			}
		}
	}
	freq := make([]int, len(clusters))
	for _, c := range clusters {
		freq[c] += 1
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freq)))
	return freq[0] * freq[1] * freq[2]
}

func dist(a []int, b []int) int {
	var x = a[0] - b[0]
	var y = a[1] - b[1]
	var z = a[2] - b[2]
	return x*x + y*y + z*z
}
