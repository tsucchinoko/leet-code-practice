package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	const INF = int(^uint(0) >> 1)
	small, mid := INF, INF

	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	tests := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 1, 5, 0, 4, 6},
		{2, 2, 2, 2},
		{1, 1, 2, 2, 3},
	}
	for _, t := range tests {
		fmt.Println(t, "=>", increasingTriplet(t))
	}
}
