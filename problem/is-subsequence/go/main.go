package main

import "fmt"

func isSubsequence(s string, t string) bool {
	si, ti := 0, 0
	sn, tn := len(s), len(t)
	for si < sn && ti < tn {
		if s[si] == t[ti] {
			si++
		}
		ti++
	}

	return si == sn
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
}
