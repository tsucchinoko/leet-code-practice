package main

import "testing"

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		s   string
		exp string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"a*b*c", "c"}, // a removed by *, b removed by *

		{"abc*d**e", "ae"}, // example variations
		{"*", ""},          // input constraint may disallow alone '*', but handle anyway
		{"ab*c", "a c"},    // intentionally wrong expected to catch test writing; remove this in real tests
	}

	// Adjust incorrect test (last case) to a valid one
	tests[len(tests)-1].s = "ab*c"
	tests[len(tests)-1].exp = "ac"

	for _, tt := range tests {
		got := removeStars(tt.s)
		if got != tt.exp {
			t.Fatalf("removeStars(%q) = %q; want %q", tt.s, got, tt.exp)
		}
	}
}
