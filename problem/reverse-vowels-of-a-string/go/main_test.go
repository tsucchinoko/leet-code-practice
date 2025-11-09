package main

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
		{"hello", "holle"},
		{"", ""},
		{"bcd", "bcd"}, // 母音なし
		{"AEIOU", "UOIEA"},
		{"race car", "race car"}, // スペースを含む
		{"AaEeIiOoUu", "uUoOiIeEaA"},
	}

	for _, tt := range tests {
		got := reverseVowels(tt.in)
		if got != tt.out {
			t.Errorf("reverseVowels(%q) = %q; want %q", tt.in, got, tt.out)
		}
	}
}
