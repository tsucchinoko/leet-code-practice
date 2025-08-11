package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		got := maxProfit(tt.prices)
		if got != tt.want {
			t.Errorf("MaxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
		}
	}
}

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		got := maxProfit2(tt.prices)
		if got != tt.want {
			t.Errorf("MaxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
		}
	}
}
