package main

// majorityElement returns the majority element in nums.
// Assumes a majority element always exists.
func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for _, x := range nums {
		if count == 0 {
			candidate = x
			count = 1
			continue
		}
		if x == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
