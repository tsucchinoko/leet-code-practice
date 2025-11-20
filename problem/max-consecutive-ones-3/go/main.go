package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left := 0
	zeroCount := 0
	maxLen := 0

	for right := range nums {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		if cur := right - left + 1; cur > maxLen {
			maxLen = cur
		}
	}
	return maxLen
}

func main() {
	samples := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
	}
	for _, sample := range samples {
		fmt.Printf("got: %d, result: %d\n", sample.nums, longestOnes(sample.nums, sample.k))
	}
}
