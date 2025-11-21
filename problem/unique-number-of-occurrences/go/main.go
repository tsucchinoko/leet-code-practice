package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	// 値→出現回数
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	// 出現回数の集合で重複を検査
	seen := make(map[int]struct{})
	for _, count := range freq {
		if _, ok := seen[count]; ok {
			return false
		}
		seen[count] = struct{}{}
	}

	return true
}

func main() {
	examples := [][]int{
		{1, 2, 2, 1, 1, 3},
		{1, 2},
		{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
	}
	for _, ex := range examples {
		fmt.Println(ex, "->", uniqueOccurrences(ex))
	}
}
