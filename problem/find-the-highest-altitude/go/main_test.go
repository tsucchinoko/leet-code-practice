package main

import "testing"

func TestHighestAltitude(t *testing.T) {
	tests := []struct {
		gain []int
		want int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
		{[]int{1, 2, 3, 4, 5}, 15},     // monotonically increasing
		{[]int{-1, -2, -3, -4, -5}, 0}, // monotonically decreasing, start is max
		{[]int{0, 0, 0, 0}, 0},         // all zeros
		{[]int{5, -2, 3, -1, 2}, 7},    // mixed
	}

	for _, tt := range tests {
		got := largestAltitude(tt.gain)
		if got != tt.want {
			t.Fatalf("highestAltitude(%v) = %d; want %d", tt.gain, got, tt.want)
		}
	}
}
