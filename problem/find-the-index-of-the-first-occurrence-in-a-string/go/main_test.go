package main

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"sadbutsad", "sad", 0},
		{"sadbutsad", "but", 3},
		{"leetcode", "leeto", -1},
		{"a", "a", 0},
		{"a", "b", -1},
		{"aaa", "aa", 0}, // first occurrence
		{"mississippi", "issip", 4},
		{"", "", 0}, // although constraints forbid empty, be robust
		{"abc", "", 0},
		{"", "a", -1},
	}

	for _, tt := range tests {
		got := strStr(tt.haystack, tt.needle)
		if got != tt.want {
			t.Errorf("strStr(%q, %q) = %d; want %d", tt.haystack, tt.needle, got, tt.want)
		}
	}
}
