package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interview", "interrupt", "integrate"}, "inte"},
		{[]string{"apple"}, "apple"},
		{[]string{}, ""},
		{[]string{"a", "a", "a"}, "a"},
		{[]string{"", ""}, ""},
		{[]string{"abc", "", "abcd"}, ""},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.input)
		if result != test.expected {
			t.Errorf("Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
		}
	}
}
