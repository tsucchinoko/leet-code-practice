package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if n == 0 || k == 0 {
		return 0.0
	}

	var sum int64 = 0
	for i := range nums[:k] {
		sum += int64(nums[i])
	}
	maxSum := sum

	for i := k; i < n; i++ {
		sum += int64(nums[i])
		sum -= int64(nums[i-k])
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)

}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
