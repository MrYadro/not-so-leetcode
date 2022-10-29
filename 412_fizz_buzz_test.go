package leetcode

import (
	"reflect"
	"testing"
)

type fizzBuzzTest struct {
	arg      int
	expected []string
}

var fizzBuzzTests = []fizzBuzzTest{
	{3, []string{"1", "2", "Fizz"}},
	{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
	{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
}

func TestFizzBuzz(t *testing.T) {

	for _, test := range fizzBuzzTests {
		if output := FizzBuzz(test.arg); reflect.DeepEqual(output, test.expected) != true {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
