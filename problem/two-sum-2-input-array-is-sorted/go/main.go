package main

import "fmt"

// twoSum returns 1-based indices [index1, index2] such that numbers[index1-1] + numbers[index2-1] == target.
// Assumes numbers is sorted in non-decreasing order and exactly one solution exists.
func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	// As per problem statement, there is always exactly one solution.
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
