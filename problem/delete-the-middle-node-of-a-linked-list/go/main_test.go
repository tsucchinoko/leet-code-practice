package main

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
		{[]int{2, 1}, []int{2}},
		{[]int{1}, []int{}},
		{[]int{1, 2, 3}, []int{1, 3}}, // n=3 middle index=1
		{[]int{1, 2}, []int{1}},       // n=2 middle index=1
	}

	for _, tt := range tests {
		head := buildList(tt.in)
		gotHead := deleteMiddle(head)
		got := toSlice(gotHead)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("deleteMiddle(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}
