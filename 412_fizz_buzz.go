package leetcode

// https://leetcode.com/problems/fizz-buzz/
// Runtime: 5 ms
// Memory Usage: 3.5 MB

import "strconv"

func FizzBuzz(n int) []string {
	m := make([]string, n)
	for i := 0; i < n; i++ {
		ip1 := i + 1
		if ip1%3 == 0 && ip1%5 == 0 {
			m[i] = "FizzBuzz"
		} else if ip1%3 == 0 {
			m[i] = "Fizz"
		} else if ip1%5 == 0 {
			m[i] = "Buzz"
		} else {
			m[i] = strconv.Itoa(ip1)
		}
	}
	return m
}
