package leetcode

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// Runtime: 0 ms
// Memory Usage: 1.9 MB

import "math/bits"

func NumberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	unum := uint(num)
	return bits.Len(unum) - 1 + bits.OnesCount(unum)
}
