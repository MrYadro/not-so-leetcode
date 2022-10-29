package leetcode

// https://leetcode.com/problems/fizz-buzz/
// Runtime: 6 ms
// Memory Usage: 3.4 MB

import "strconv"

func FizzBuzz(n int) []string {
	m := make([]string, n)
	for i := 0; i < n; i++ {
		ip1 := i + 1
		db3 := ip1%3 == 0
		db5 := ip1%5 == 0
		if db3 && db5 {
			m[i] = "FizzBuzz"
		} else if db3 {
			m[i] = "Fizz"
		} else if db5 {
			m[i] = "Buzz"
		} else {
			m[i] = strconv.Itoa(ip1)
		}
	}
	return m
}
