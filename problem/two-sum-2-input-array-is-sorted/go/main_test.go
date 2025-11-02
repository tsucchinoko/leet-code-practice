package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 4, 9}, 8, []int{4, 5}},     // 4 + 4 = 8
		{[]int{-5, -3, 0, 2, 4, 10}, -1, []int{1, 5}}, // -5 + 4 = -1
	}

	for _, tc := range tests {
		got := twoSum(tc.numbers, tc.target)
		if got == nil || len(got) != 2 || got[0] != tc.want[0] || got[1] != tc.want[1] {
			t.Errorf("twoSum(%v, %d) = %v; want %v", tc.numbers, tc.target, got, tc.want)
		}
	}
}
