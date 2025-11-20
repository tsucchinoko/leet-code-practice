package main

import "fmt"

func longestSubarray(nums []int) int {
	l := 0
	zeros := 0
	maxLen := 0

	for r := range nums {
		if nums[r] == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		// ウィンドウ [l, r] は 0 を最大1つ含む
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen - 1
}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                // 3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) // 5
	fmt.Println(longestSubarray([]int{1, 1, 1}))                   // 2
}
