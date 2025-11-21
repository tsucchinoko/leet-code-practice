package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var cnt1 [26]int
	var cnt2 [26]int
	for i := 0; i < len(word1); i++ {
		cnt1[word1[i]-'a']++
	}
	for i := 0; i < len(word2); i++ {
		cnt2[word2[i]-'a']++
	}

	// 出現文字集合が同じか確認
	for i := range cnt1 {
		if (cnt1[i] == 0) != (cnt2[i] == 0) {
			return false
		}
	}

	// 出現頻度のマルチセットを比較（非ゼロのみ）
	freqs1 := make([]int, 0, 26)
	freqs2 := make([]int, 0, 26)
	for i := range cnt1 {
		if cnt1[i] > 0 {
			freqs1 = append(freqs1, cnt1[i])
		}
		if cnt2[i] > 0 {
			freqs2 = append(freqs2, cnt2[i])
		}
	}
	sort.Ints(freqs1)
	sort.Ints(freqs2)
	if len(freqs1) != len(freqs2) {
		return false
	}
	for i := range freqs1 {
		if freqs1[i] != freqs2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))       // true
	fmt.Println(closeStrings("a", "aa"))          // false
	fmt.Println(closeStrings("cabbba", "abbccc")) // true
	fmt.Println(closeStrings("cabbba", "aabbss")) // false
}
