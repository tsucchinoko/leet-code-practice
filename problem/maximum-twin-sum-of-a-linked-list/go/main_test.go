package main

import "testing"

func TestPairSumExamples(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{5, 4, 2, 1}, 6},
		{[]int{4, 2, 2, 3}, 7},
		{[]int{1, 100000}, 100001},
		{[]int{1, 2, 3, 4, 5, 6}, 7}, // pairs: (1+6)=7,(2+5)=7,(3+4)=7
		{[]int{10, 20, 30, 40}, 50},  // (10+40)=50,(20+30)=50
	}

	for _, tt := range tests {
		head := buildList(tt.vals)
		got := pairSum(head)
		if got != tt.want {
			t.Fatalf("vals=%v: got %d, want %d", tt.vals, got, tt.want)
		}
	}
}
