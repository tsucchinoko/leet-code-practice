package main

import (
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{in: []int{1, 2, 3, 4, 5}, out: []int{1, 3, 5, 2, 4}},
		{in: []int{2, 1, 3, 5, 6, 4, 7}, out: []int{2, 3, 6, 7, 1, 5, 4}},
		{in: []int{}, out: []int{}},
		{in: []int{1}, out: []int{1}},
		{in: []int{1, 2}, out: []int{1, 2}},
		{in: []int{1, 2, 3}, out: []int{1, 3, 2}},
	}

	for _, c := range cases {
		got := oddEvenList(buildList(c.in))
		gotSlice := toSlice(got)
		if !reflect.DeepEqual(gotSlice, c.out) {
			t.Fatalf("oddEvenList(%v) = %v; want %v", c.in, gotSlice, c.out)
		}
	}
}
