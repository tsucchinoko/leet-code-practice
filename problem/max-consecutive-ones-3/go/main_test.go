package main

import (
	"reflect"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{1, 1, 1, 1}, 0, 4},
		{[]int{0, 0, 0, 0}, 2, 2},
		{[]int{0, 1, 0, 1, 1, 0, 1}, 1, 4},
		{[]int{1}, 0, 1},
		{[]int{0}, 0, 0},
		{[]int{0}, 1, 1},
	}

	for _, tt := range tests {
		got := longestOnes(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("longestOnes(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
