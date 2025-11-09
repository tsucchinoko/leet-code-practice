package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	b := []byte(s)
	left, right := 0, len(s)-1

	for left < right {
		// 左を母音に移動
		for left < right && !vowels[b[left]] {
			left++
		}
		// 右を母音に移動
		for left < right && !vowels[b[right]] {
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
