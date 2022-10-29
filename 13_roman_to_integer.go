package leetcode

// Runtime: 7 ms
// Memory Usage: 3.6 MB

import (
	"strings"
)

func RomanToInt(s string) int {
	ss := strings.Split(s, "")
	rti := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sum := 0
	cur := rti[ss[len(ss)-1]]
	for i := len(ss) - 1; i >= 0; i-- {
		next := rti[ss[i]]
		if cur <= next {
			sum += next
		} else {
			sum -= next
		}
		cur = rti[ss[i]]
	}
	return sum
}
