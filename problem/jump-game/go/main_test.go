package main

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want bool
	}{
		{"example1", []int{2, 3, 1, 1, 4}, true},
		{"example2", []int{3, 2, 1, 0, 4}, false},
		{"single", []int{0}, true},
		{"cannotMove", []int{0, 1}, false},
		{"longJump", []int{4, 0, 0, 0, 0}, true},
	}

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := canJump(tc.nums)
			if got != tc.want {
				t.Fatalf("canJump(%v) = %v; want %v", tc.nums, got, tc.want)
			}
		})
	}
}
