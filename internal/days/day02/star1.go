package day02

import (
	"strconv"
	"strings"
)

func Star1(input string) any {
	var lines = strings.Split(input, ",")
	var count = 0
	for _, set := range lines {
		s := strings.Split(set, "-")
		var start, _ = strconv.Atoi(s[0])
		var end, _ = strconv.Atoi(s[1])
		for i := start; i <= end; i++ {
			if repeating(i) {
				count += i
			}
		}
	}
	return count
}

func repeating(num int) bool {
	var len = digits(num)
	if len%2 == 0 {
		var pow = 1
		for range len / 2 {
			pow *= 10
		}
		var first = num % pow
		for i := num / pow; i > 0; i = i / pow {
			if first != i%pow {
				return false
			}
		}
		return true
	}

	return false
}

func digits(i int) int {
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}
