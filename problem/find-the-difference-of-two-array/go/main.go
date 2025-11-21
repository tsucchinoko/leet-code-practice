package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}

	res1 := make([]int, 0)
	for v := range set1 {
		if _, ok := set2[v]; !ok {
			res1 = append(res1, v)
		}
	}

	res2 := make([]int, 0)
	for v := range set2 {
		if _, ok := set1[v]; !ok {
			res2 = append(res2, v)
		}
	}

	return [][]int{res1, res2}

}

func main() {
	// 簡単な実行例
	a := []int{1, 2, 3}
	b := []int{2, 4, 6}
	fmt.Println(findDifference(a, b)) // 例: [[1 3] [4 6]] （順序は異なる可能性あり）
}
