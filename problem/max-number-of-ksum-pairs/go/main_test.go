package main

import "testing"

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{1, 1, 1, 1}, 2, 2},
		{[]int{2, 2, 2, 3, 3, 3}, 5, 3},
		{[]int{1}, 2, 0},
		{[]int{1, 4, 1, 4, 2, 3}, 5, 3},
		{[]int{1000000000, 1}, 1000000001, 1},
	}

	for _, tt := range tests {
		got := maxOperations(tt.nums, tt.k)
		if got != tt.want {
			t.Fatalf("maxOperations(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
