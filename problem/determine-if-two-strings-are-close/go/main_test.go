package main

import "testing"

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		w1   string
		w2   string
		want bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"abbzzca", "babzzcz", false}, // 異なる出現集合
		{"aaabbb", "ababab", true},    // 同じ頻度分布・集合
		{"abc", "def", false},         // 使われる文字集合が違う
		{"", "", true},                // 空文字（制約では長さ>=1だが保険）
	}

	for _, tt := range tests {
		got := closeStrings(tt.w1, tt.w2)
		if got != tt.want {
			t.Errorf("closeStrings(%q, %q) = %v; want %v", tt.w1, tt.w2, got, tt.want)
		}
	}
}
