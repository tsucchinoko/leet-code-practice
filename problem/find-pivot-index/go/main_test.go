package main

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
		{[]int{0}, 0},
		{[]int{-1, -1, -1, 0, 1, 1}, 0},
		{[]int{1, -1, 0}, 2},
		{[]int{1, -1, 1, -1, 1, -1}, -1}, // all prefixes match at 0
	}

	for _, tt := range tests {
		got := pivotIndex(tt.nums)
		if got != tt.want {
			t.Fatalf("pivotIndex(%v) = %d; want %d", tt.nums, got, tt.want)
		}
	}
}
