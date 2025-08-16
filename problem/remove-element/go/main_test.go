package main

import (
	"testing"
)

// func TestRemoveElement(t *testing.T) {
// 	tests := []struct {
// 		nums   []int
// 		val    int
// 		expect int
// 	}{
// 		{[]int{3, 2, 2, 3}, 3, 2},
// 		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
// 	}

// 	for _, test := range tests {
// 		got := removeElement(test.nums, test.val)
// 		if got != test.expect {
// 			t.Errorf("removeElement(%v, %d) = %d; want %d", test.nums, test.val, got, test.expect)
// 		}
// 	}
// }

// ベンチマーク: 前から詰める方
func BenchmarkRemoveElement(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i % 10 // 0〜9 の値を繰り返す
	}
	val := 5
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// コピーして使わないとテストに影響が出るので毎回コピーする
		data := make([]int, len(nums))
		copy(data, nums)
		removeElement(data, val)
	}
}

// ベンチマーク: 末尾と交換する方
func BenchmarkRemoveElement2(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i % 10
	}
	val := 5
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data := make([]int, len(nums))
		copy(data, nums)
		removeElement2(data, val)
	}
}
