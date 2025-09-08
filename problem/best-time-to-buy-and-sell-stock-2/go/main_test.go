package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1}, 0},
		{[]int{2, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{3, 3, 3}, 0},
		{[]int{1, 3, 2, 5, 4, 6}, 7}, // 1->3 (2) + 2->5 (3) + 4->6 (2) = 7
	}

	for _, tt := range tests {
		got := maxProfit(tt.prices)
		if got != tt.want {
			t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
		}
	}
}
