package main

import "fmt"

// strStr returns the index of the first occurrence of needle in haystack, or -1 if not present.
func strStr(haystack string, needle string) int {
	hn, nn := len(haystack), len(needle)
	if nn == 0 {
		return 0
	}
	if hn < nn {
		return -1
	}

	// Sliding window native search
	// i: index of the first character of the current window = offset
	// j: index of the first character of the current substring = parts of the target needle
	for i := 0; i <= hn-nn; i++ {
		j := 0
		// check offset + parts of the target needle
		for j < nn && haystack[i+j] == needle[j] {
			j++
		}
		// needle is all matched
		if j == nn {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))  // 0
	fmt.Println(strStr("leetcode", "leeto")) // -1
}
