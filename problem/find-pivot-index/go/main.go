package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == total-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

func main() {
	examples := [][]int{
		{1, 7, 3, 6, 5, 6},
		{1, 2, 3},
		{2, 1, -1},
	}
	for _, ex := range examples {
		fmt.Println(pivotIndex(ex))
	}
}
