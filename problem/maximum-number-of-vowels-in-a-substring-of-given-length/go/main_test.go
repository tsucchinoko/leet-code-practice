package main

import "testing"

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"rhythms", 4, 0},   // 母音なし
		{"a", 1, 1},         // 最小ケース
		{"abcde", 5, 2},     // 全長ウィンドウ
		{"abecidofu", 3, 2}, // 複数母音の分布
	}

	for _, tt := range tests {
		got := maxVowels(tt.s, tt.k)
		if got != tt.want {
			t.Fatalf("maxVowels(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
		}
	}
}
