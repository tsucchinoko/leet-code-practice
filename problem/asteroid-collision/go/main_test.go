package main

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{3, 5, -6, 2, -1, 4}, []int{-6, 2, 4}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}}, // no collisions (same directions or separated)
		{[]int{1, -2, -2, -2}, []int{-2, -2, -2}},
		{[]int{2, 1, -1, -2}, []int{}},
	}

	for _, tc := range tests {
		got := asteroidCollision(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Fatalf("input %v: expected %v, got %v", tc.in, tc.out, got)
		}
	}
}
