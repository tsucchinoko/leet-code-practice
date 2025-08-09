package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	want := []int{0, 1}
	got := twoSum(nums, target)
	if len(got) != len(want) {
		t.Errorf("length mismatch")
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("value mismatch at index %d: got %v, want %v", i, got, want)
		}
	}

}
