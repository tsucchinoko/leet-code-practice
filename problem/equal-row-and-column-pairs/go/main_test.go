package main

import "testing"

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			want: 1,
		},
		{
			grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			want: 3,
		},
		{
			grid: [][]int{{1}},
			want: 1,
		},
		{
			grid: [][]int{{1, 2}, {2, 1}},
			want: 0,
		},
		{
			grid: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			want: 9, // 全ての行と列が [1,1,1] なので 3*3 = 9
		},
	}

	for _, tt := range tests {
		if got := equalPairs(tt.grid); got != tt.want {
			t.Fatalf("equalPairs(%v) = %d; want %d", tt.grid, got, tt.want)
		}
	}
}
