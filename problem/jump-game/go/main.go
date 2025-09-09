package main

// canJump returns true if you can reach the last index of nums.
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	maxReach := 0
	last := len(nums) - 1
	for i := 0; i <= maxReach && i < len(nums); i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		if maxReach >= last {
			return true
		}
	}
	return false
}
