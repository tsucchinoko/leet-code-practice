package main

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
		{"AAAA", "AA", "AA"},
		{"ABCABCABC", "ABCABC", "ABC"},
		{"XYZ", "XYZXYZ", "XYZ"},
		{"ABAB", "AB", "AB"},
		{"ABAB", "ABA", ""}, // 長さの観点で合わない例
		{"", "", ""},        // 仕様上長さ>=1だが堅牢性のためのテスト
	}
	for _, tt := range tests {
		got := gcdOfStrings(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("gcdOfStrings(%q, %q) = %q; want %q", tt.a, tt.b, got, tt.want)
		}
	}
}
