package main

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
		{[]int{0, 0, 0, 0}, 2, 0.0},
		{[]int{-1, -2, -3, -4}, 2, -1.5},
		{[]int{1, 2, 3, 4, 5}, 5, 3.0},
		{[]int{1, 2, 3, 4, 5}, 1, 5.0},
	}

	eps := 1e-5
	for _, tc := range tests {
		got := findMaxAverage(tc.nums, tc.k)
		if !almostEqual(got, tc.want, eps) {
			t.Errorf("findMaxAverage(%v, %d) = %v; want %v", tc.nums, tc.k, got, tc.want)
		}
	}
}
