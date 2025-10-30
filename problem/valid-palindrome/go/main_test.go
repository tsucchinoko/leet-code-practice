package main

import "testing"

func TestIsPalindromeTwoPointers(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"", true},
		{"0P", false},
		{"Able was I, ere I saw Elba", true},
		{"No 'x' in Nixon", true},
		{"Madam, I'm Adam", true},
		{"12321", true},
		{"1231", false},
	}

	for _, tc := range tests {
		got := isPalindrome(tc.input)
		if got != tc.want {
			t.Errorf("isPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
