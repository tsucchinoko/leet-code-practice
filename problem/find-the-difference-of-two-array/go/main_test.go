package main

import (
	"reflect"
	"sort"
	"testing"
)

// helper to compare two slices of slices irrespective of inner order and outer order
func equalResult(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		sa := append([]int(nil), a[i]...)
		sb := append([]int(nil), b[i]...)
		sort.Ints(sa)
		sort.Ints(sb)
		if !reflect.DeepEqual(sa, sb) {
			return false
		}
	}
	return true
}

func TestFindDifference(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
		{[]int{1}, []int{1}, [][]int{{}, {}}},
		{[]int{0, -1, -1, 2}, []int{-1, 3}, [][]int{{0, 2}, {3}}},
		{[]int{1000, -1000}, []int{500, -1000}, [][]int{{1000}, {500}}},
	}

	for _, tc := range tests {
		got := findDifference(tc.nums1, tc.nums2)
		if !equalResult(got, tc.want) {
			t.Errorf("FindDifference(%v, %v) = %v; want %v", tc.nums1, tc.nums2, got, tc.want)
		}
	}
}
