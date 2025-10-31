package main

import "testing"

func TestIsSubsequenceSimple(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "ahbgdc", true},
		{"a", "", false},
		{"", "", true},
		{"ace", "abcde", true},
		{"aec", "abcde", false},
	}
	for _, tc := range tests {
		got := isSubsequence(tc.s, tc.t)
		if got != tc.want {
			t.Fatalf("isSubsequence(%q, %q) = %v; want %v", tc.s, tc.t, got, tc.want)
		}
	}
}

// func TestSubsequenceChecker(t *testing.T) {
// 	tstr := "ahbgdc"
// 	sc := NewSubsequenceChecker(tstr)

// 	tests := []struct {
// 		s    string
// 		want bool
// 	}{
// 		{"abc", true},
// 		{"axc", false},
// 		{"", true},
// 		{"ahbgdc", true},
// 		{"ahbgdcz", false},
// 		{"hb", true},
// 	}

// 	for _, tc := range tests {
// 		got := sc.IsSubsequence(tc.s)
// 		if got != tc.want {
// 			t.Fatalf("NewSubsequenceChecker(%q).IsSubsequence(%q) = %v; want %v", tstr, tc.s, got, tc.want)
// 		}
// 	}
// }

// // Benchmark-style test for many queries simulated (not using testing.B).
// func TestManyQueries(t *testing.T) {
// 	tstr := "abacbabcdelmnopqrsabctuvwxyzab" // sample t
// 	sc := NewSubsequenceChecker(tstr)

// 	// simulate many small s queries
// 	queries := []string{"abc", "abct", "mnop", "zz", "ab", "qrst", "abctu"}
// 	for _, q := range queries {
// 		_ = sc.IsSubsequence(q) // ensure no panic and runs fast
// 	}
// }
