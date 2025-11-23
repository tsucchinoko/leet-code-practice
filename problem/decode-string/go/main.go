package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	var strStack []string
	var countStack []int
	cur := "" // 現在の部分文字列を直接保持
	num := 0

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			countStack = append(countStack, num)
			strStack = append(strStack, cur)
			num = 0
			cur = ""
		} else if ch == ']' {
			k := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			cur = prev + strings.Repeat(cur, k)
		} else {
			cur += string(ch)
		}
	}
	return cur
}

func main() {
	// 簡単なデモ
	examples := []string{"3[a]2[bc]", "3[a2[c]]", "2[abc]3[cd]ef"}
	for _, e := range examples {
		fmt.Println(e, "->", decodeString(e))
	}
}
