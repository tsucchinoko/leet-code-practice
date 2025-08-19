package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, test := range tests {
		actual := removeDuplicates(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("removeDuplicates(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}
