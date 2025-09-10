package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestJump_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{2, 3, 1, 1, 4}, 2},
		{"example2", []int{2, 3, 0, 1, 4}, 2},
		{"single_zero", []int{0}, 0},
		{"all_ones", []int{1, 1, 1, 1}, 3},
		// corrected expected value: from 1 -> 2 because nums[0]=4 reaches index 4, not the last index 6
		{"can_reach_in_two", []int{4, 1, 1, 3, 1, 1, 1}, 2},
		{"mixed_zeros", []int{2, 0, 2, 0, 1}, 2},
		{"large_step_at_start", []int{5, 0, 0, 0, 0, 0}, 1},
		{"two_elements", []int{0, 1}, 1},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := jump(tt.nums)
			if got != tt.want {
				t.Errorf("jump(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestJump_RandomizedReachableCases(t *testing.T) {
	// Use a local rand.Source and rand.Rand to avoid deprecated global Seed usage
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	const trials = 100
	const maxLen = 50

	// create an index slice to range over (modern range-int style)
	idxs := make([]int, trials)
	for i := range idxs {
		idxs[i] = i
	}

	for range idxs {
		n := r.Intn(maxLen-1) + 2 // length in [2, maxLen]
		nums := make([]int, n)

		// create a helper slice to range from 0 to n-2
		positions := make([]int, n-1)
		for i := range positions {
			positions[i] = i
		}

		for _, i := range positions {
			maxJump := n - 1 - i
			if r.Intn(5) == 0 {
				nums[i] = maxJump // make this index reach the end
			} else {
				nums[i] = r.Intn(maxJump + 1)
			}
		}
		nums[n-1] = 0

		got := jump(nums)
		if got < 0 {
			t.Errorf("unexpected negative jumps for nums=%v", nums)
		}
		if got > n-1 {
			t.Errorf("too many jumps (%d) for nums=%v (len=%d)", got, nums, n)
		}
	}
}
