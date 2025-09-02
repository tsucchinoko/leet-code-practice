package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1}, 0, []int{1}},
		{[]int{1, 2}, 3, []int{2, 1}},       // k > n
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}}, // k == n
		{[]int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}},
	}

	for _, tc := range tests {
		// コピーして関数を呼ぶ（破壊的なので）
		arr := make([]int, len(tc.nums))
		copy(arr, tc.nums)
		rotate(arr, tc.k)
		if !reflect.DeepEqual(arr, tc.expected) {
			t.Errorf("rotate(%v, %d) = %v; expected %v", tc.nums, tc.k, arr, tc.expected)
		}
	}
}
