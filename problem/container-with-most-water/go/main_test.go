package main

import "testing"

func TestMaxAreaExamples(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},   // both ends
		{[]int{1, 2, 1}, 2},          // middle not best
		{[]int{0, 0, 0, 0}, 0},       // zeros
		{[]int{10000, 10000}, 10000}, // large values
		{[]int{1, 2, 4, 3}, 4},
	}

	for _, tt := range tests {
		got := maxArea(tt.height)
		if got != tt.want {
			t.Fatalf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
		}
	}
}
