package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if visited_index, ok := m[diff]; ok {
			return []int{visited_index, i}
		}
		m[num] = i
	}
	panic("solution was not found!")
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("TwoSum result:", result)
}
