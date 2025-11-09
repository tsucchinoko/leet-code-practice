package main

import (
	"fmt"
)

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}
	b := []byte(s)
	left, right := 0, len(s)-1

	for left < right {
		// 左を母音に移動
		for left < right && !isVowel(b[left]) {
			left++
		}
		// 右を母音に移動
		for left < right && !isVowel(b[right]) {
			right--
		}
		// 左右の母音を交換
		if left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}

	return string(b)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm")) // AceCreIm
	fmt.Println(reverseVowels("leetcode")) // leotcede
}
