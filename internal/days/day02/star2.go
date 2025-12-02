package day02

import (
	"strconv"
	"strings"
)

func Star2(input string) any {
	var lines = strings.Split(input, ",")
	var count = 0
	for _, set := range lines {
		s := strings.Split(set, "-")
		var start, _ = strconv.Atoi(s[0])
		var end, _ = strconv.Atoi(s[1])
		for i := start; i <= end; i++ {
			if repeating2(i) {
				count += i
			}
		}
	}
	return count
}

func repeating2(num int) bool {
	var len = digits(num)
	for times := 2; times <= len; times++ {
		if len%times == 0 {
			var pow = 1
			for range len / times {
				pow *= 10
			}
			var first = num % pow
			var fail = false
			for i := num / pow; i > 0; i = i / pow {
				if first != i%pow {
					fail = true
				}
			}
			if !fail {
				return true
			}
		}
	}
	return false
}
