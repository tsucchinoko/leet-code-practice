package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// prefixが各strs[i]の接頭辞になっていなければ末尾を1文字ずつ削る
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

// 指定範囲[left, right]の長さの部分文字列が全ての文字列で共通かチェック
func allContainPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < length || strs[i][:length] != prefix {
			return false
		}
	}
	return true
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 最短文字列の長さを求める
	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	lo, hi := 0, minLen
	for lo < hi {
		mid := (lo + hi + 1) / 2 // +1を入れることでloが進まない無限ループを防止
		if allContainPrefix(strs, mid) {
			lo = mid // midの長さで共通接頭辞があるなら伸ばす
		} else {
			hi = mid - 1 // なければ縮める
		}
	}

	return strs[0][:lo]
}
