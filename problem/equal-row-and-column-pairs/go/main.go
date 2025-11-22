package main

import (
	"fmt"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rowCount := make(map[string]int, n)

	// 行をキー化してカウント
	for _, row := range grid {
		parts := make([]string, n)
		for i, v := range row {
			parts[i] = fmt.Sprintf("%d", v)
		}
		key := strings.Join(parts, ",")
		rowCount[key]++
	}

	// 各列をキー化して rowCount に存在すれば加算
	res := 0
	for j := range grid[0] {
		parts := make([]string, n)
		for i, row := range grid {
			parts[i] = fmt.Sprintf("%d", row[j])
		}
		key := strings.Join(parts, ",")
		if c, ok := rowCount[key]; ok {
			res += c
		}
	}
	return res
}

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))

}
