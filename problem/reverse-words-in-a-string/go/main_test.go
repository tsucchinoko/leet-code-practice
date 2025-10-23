package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"single", "single"},
		{"   leading", "leading"},
		{"trailing   ", "trailing"},
		{"  multiple   spaces  between words ", "words between spaces multiple"},
		{"  a  b ", "b a"},
	}

	for _, tt := range tests {
		got := reverseWords(tt.in)
		if got != tt.want {
			t.Errorf("reverseWords(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
