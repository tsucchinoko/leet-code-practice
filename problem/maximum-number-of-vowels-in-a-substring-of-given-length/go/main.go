package main

import "fmt"

func maxVowels(s string, k int) int {
	isVowels := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}
	n := len(s)
	if k > n {
		k = n
	}
	// 初期ウィンドウ
	count := 0
	for i := 0; i < k; i++ {
		if isVowels(s[i]) {
			count++
		}
	}
	maxCount := count

	// スライドウィンドウ
	for i := k; i < n; i++ {
		if isVowels(s[i-k]) {
			count--
		}
		if isVowels(s[i]) {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func main() {
	samples := []struct {
		s string
		k int
	}{
		{"abciiidef", 3},
		{"aeiou", 2},
		{"leetcode", 3},
	}
	for _, sample := range samples {
		fmt.Printf("called with %v %d, result: {%d}\n", sample.s, sample.k, maxVowels(sample.s, sample.k))
	}
}
