package leetcode

import "testing"

type romanToIntegerTest struct {
	arg      string
	expected int
}

var romanToIntegerTests = []romanToIntegerTest{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInteger(t *testing.T) {

	for _, test := range romanToIntegerTests {
		if output := RomanToInt(test.arg); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
