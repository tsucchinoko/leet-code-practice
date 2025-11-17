package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 0, 2, 0, 0, 3}, []int{1, 2, 3, 0, 0, 0}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 1, 2, 0, 3, 0, 4}, []int{1, 2, 3, 4, 0, 0, 0}},
	}

	for _, tt := range tests {
		input := make([]int, len(tt.input))
		copy(input, tt.input)
		moveZeroes(input)
		if !reflect.DeepEqual(input, tt.expect) {
			t.Errorf("moveZeroes(%v) = %v; want %v", tt.input, input, tt.expect)
		}
	}
}
