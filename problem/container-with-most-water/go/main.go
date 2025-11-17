package main

import (
	"fmt"
)

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	left, right := 0, n-1
	maxArea := 0
	for left < right {
		h := min(height[left], height[right])
		area := (right - left) * h
		if area > maxArea {
			maxArea = area
		}
		// move the smaller height pointer inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	examples := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	for _, ex := range examples {
		fmt.Printf("height = %v -> maxArea = %d\n", ex, maxArea(ex))
	}
}
