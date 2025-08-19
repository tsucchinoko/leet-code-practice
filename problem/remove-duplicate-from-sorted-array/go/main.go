package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// すでにソート済みの前提のため、異なる場合は新しい数字
		if nums[fast] != nums[slow] {
			// 新しい数字と重複した数字を入れ替える
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
