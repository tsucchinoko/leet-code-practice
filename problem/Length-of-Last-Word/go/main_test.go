package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
		{"a ", 1},
		{" abc", 3},
		{"abc  def   ", 3}, // last word "def"
		{"   single", 6},
		{"space at end ", 3}, // "end"
	}

	for _, tt := range tests {
		got := lengthOfLastWord(tt.s)
		if got != tt.want {
			t.Errorf("lengthOfLastWord(%q) = %d; want %d", tt.s, got, tt.want)
		}
	}
}
