package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
		// 境界ケース
		{[]int{5}, true},                      // 1要素は常に true
		{[]int{2, 2, 2, 2}, true},             // 1つの値のみ -> 出現回数集合は {4}
		{[]int{1, 1, 2, 2, 3, 3}, false},      // すべて2回ずつ -> 重複
		{[]int{-1, -1, 0, 0, 0}, true},        // -1:2回, 0:3回 -> 出現は {2,3} -> true
		{[]int{-1, -1, 0, 0, 0, 7, 7}, false}, // -1:2,0:3,7:2 -> 2が重複 -> false
	}

	for _, tt := range tests {
		got := uniqueOccurrences(tt.arr)
		if got != tt.want {
			t.Errorf("uniqueOccurrences(%v) = %v; want %v", tt.arr, got, tt.want)
		}
	}
}
