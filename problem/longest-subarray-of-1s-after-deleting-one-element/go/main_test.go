package main

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
		{[]int{0, 0, 0}, 0},
		{[]int{1}, 0}, // 1要素で1を削除すると空なので0
		{[]int{0}, 0},
		{[]int{1, 0, 1, 1, 0, 1, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 0, 0, 1, 1, 1}, 3},
	}

	for _, tc := range tests {
		got := longestSubarray(tc.nums)
		if got != tc.want {
			t.Fatalf("nums=%v: got=%d want=%d", tc.nums, got, tc.want)
		}
	}
}
