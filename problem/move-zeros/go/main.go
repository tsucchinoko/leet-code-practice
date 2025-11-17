package main

import "fmt"

func moveZeroes(nums []int) {
	write := 0
	for i := range nums {
		if nums[i] != 0 {
			// i == write の時は入れ替えても同じためスキップ
			if i != write {
				nums[write] = nums[i]
			}
			write++
		}
	}
	for write < len(nums) {
		nums[write] = 0
		write++
	}
}

func main() {
	examples := [][]int{
		{0, 1, 0, 3, 12},
		{0},
		{1, 0, 2, 0, 0, 3},
		{0, 0, 0},
		{1, 2, 3},
	}

	for _, ex := range examples {
		nums := make([]int, len(ex))
		copy(nums, ex)
		moveZeroes(nums)
		fmt.Println(nums)
	}
}
