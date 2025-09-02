package main

import (
	"testing"
)

func TestMajorityElement_Basic(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{3, 2, 3}, 3},
		{"example2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"single", []int{7}, 7},
		{"all same", []int{5, 5, 5, 5, 5}, 5},
		{"majority at end", []int{1, 2, 1, 1, 3, 1, 1}, 1},
		{"large mix", []int{9, 9, 8, 7, 9, 6, 9, 9, 2, 9, 9}, 9},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := majorityElement(c.nums)
			if got != c.want {
				t.Fatalf("majorityElement(%v) = %d; want %d", c.nums, got, c.want)
			}
		})
	}
}

func TestMajorityElement_VerifyWhenNotGuaranteed(t *testing.T) {
	// This test demonstrates how to verify candidate when majority is not guaranteed.
	// For our problem the majority always exists, so majorityElement() is sufficient.
	nums := []int{1, 2, 3, 2, 2}
	candidate := majorityElement(nums)

	// verify
	count := 0
	for _, v := range nums {
		if v == candidate {
			count++
		}
	}
	if count <= len(nums)/2 {
		t.Fatalf("candidate %d is not a majority in %v", candidate, nums)
	}
}
