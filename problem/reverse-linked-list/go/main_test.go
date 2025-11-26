package main

import (
	"reflect"
	"testing"
)

func TestReverseIterative(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"one", []int{1}, []int{1}},
		{"two", []int{1, 2}, []int{2, 1}},
		{"five", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"negatives", []int{-1, 0, 3}, []int{3, 0, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.in)
			got := reverseList(head)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Fatalf("reverseIterative(%v) = %v, want %v", tt.in, listToSlice(got), tt.want)
			}
		})
	}
}

func TestReverseRecursive(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"one", []int{1}, []int{1}},
		{"two", []int{1, 2}, []int{2, 1}},
		{"five", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"negatives", []int{-1, 0, 3}, []int{3, 0, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.in)
			got := reverseRecursive(head)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Fatalf("reverseRecursive(%v) = %v, want %v", tt.in, listToSlice(got), tt.want)
			}
		})
	}
}
