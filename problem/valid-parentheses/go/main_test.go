package main

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, test := range tests {
		if isValidParentheses(test.input) != test.expected {
			t.Errorf("isValidParentheses(%q) = %v; want %v", test.input, isValidParentheses(test.input), test.expected)
		}
	}
}
