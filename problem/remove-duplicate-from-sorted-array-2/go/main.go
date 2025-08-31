package main

// removeDuplicates は与えられたソート済みスライス nums をインプレースで変更し、
// 各値を最大2回まで残して、その結果の長さ k を返す。
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	j := 2
	for i := 2; i < n; i++ {
		// nums[j-2] が結果配列内で2つ前の値
		// nums[i] が等しくなければ受け入れて書き込む
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
