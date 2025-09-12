package main

import "testing"

// バケット法
func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{"example1", []int{3, 0, 6, 1, 5}, 3},
		{"example2", []int{1, 3, 1}, 1},
		{"all_zero", []int{0, 0, 0}, 0},
		{"single_one", []int{1}, 1},
		{"single_zero", []int{0}, 0},
		{"large_values", []int{1000, 1000, 1000}, 3},
		{"mixed", []int{0, 1, 3, 5, 6}, 3},
		{"mixed2", []int{4, 4, 0, 0}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hIndex(tt.citations)
			if got != tt.want {
				t.Fatalf("hIndexCount(%v) = %d; want %d", tt.citations, got, tt.want)
			}
		})
	}
}

// 降順ソート法
func TestHIndex2(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{"example1", []int{3, 0, 6, 1, 5}, 3},
		{"example2", []int{1, 3, 1}, 1},
		{"all_zero", []int{0, 0, 0}, 0},
		{"single_one", []int{1}, 1},
		{"single_zero", []int{0}, 0},
		{"large_values", []int{1000, 1000, 1000}, 3},
		{"mixed", []int{0, 1, 3, 5, 6}, 3},
		{"mixed2", []int{4, 4, 0, 0}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hIndex2(tt.citations)
			if got != tt.want {
				t.Fatalf("hIndexSort(%v) = %d; want %d", tt.citations, got, tt.want)
			}
		})
	}
}
