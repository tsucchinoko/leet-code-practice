package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			want:         []bool{true, false, true},
		},
		{
			candies:      []int{5, 5, 5},
			extraCandies: 0,
			want:         []bool{true, true, true},
		},
		{
			candies:      []int{1, 2},
			extraCandies: 50,
			want:         []bool{true, true},
		},
	}

	for _, tc := range tests {
		got := kidsWithCandies(tc.candies, tc.extraCandies)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("candies=%v extra=%d => got=%v want=%v", tc.candies, tc.extraCandies, got, tc.want)
		}
	}
}
