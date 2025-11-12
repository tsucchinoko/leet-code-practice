package main

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
		{[]int{2, 2, 2, 2}, false},
		{[]int{1, 1, 2, 2, 3}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1, 2}, false},
		{[]int{}, false},
		{[]int{0, -1, 2, -2, 3}, true},
	}

	for _, c := range cases {
		got := increasingTriplet(c.nums)
		if got != c.want {
			t.Errorf("increasingTriplet(%v) = %v; want %v", c.nums, got, c.want)
		}
	}
}
