package leetcode

import "testing"

type ransomNoteTest struct {
	arg1     string
	arg2     string
	expected bool
}

var ransomNoteTests = []ransomNoteTest{
	{"a", "b", false},
	{"aa", "ab", false},
	{"aa", "aab", true},
}

func TestRansomNote(t *testing.T) {

	for _, test := range ransomNoteTests {
		if output := CanConstruct(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
