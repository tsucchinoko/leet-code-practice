package main

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"aba", true},
		{"abca", true},
		{"abc", false},
	}

	for _, tt := range tests {
		if got := validPalindrome(tt.s); got != tt.want {
			t.Errorf("validPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
