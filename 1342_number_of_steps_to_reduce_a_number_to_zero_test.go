package leetcode

import "testing"

type reduceNumberTest struct {
	arg      int
	expected int
}

var reduceNumberTests = []reduceNumberTest{
	{14, 6},
	{8, 4},
	{123, 12},
}

func TestReduceNumber(t *testing.T) {

	for _, test := range reduceNumberTests {
		if output := NumberOfSteps(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
