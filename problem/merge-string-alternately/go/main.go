package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte, 0, len(word1)+len(word2))
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			res = append(res, word1[i])
		}
		if i < len(word2) {
			res = append(res, word2[i])
		}
	}
	return string(res)
}

// - 内部でバッファを管理しており、繰り返しの連結で何度もメモリを確保し直すことを抑制できる。
// - `[]byte` で自分でメモリ管理をするよりもシンプルに書ける。
// - 大量の文字列を繰り返し連結する場合のパフォーマンスが良い
func mergeAlternately2(word1, word2 string) string {
	var b strings.Builder
	b.Grow(len(word1) + len(word2)) // 十分な容量を確保

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			b.WriteByte(word1[i])
		}
		if i < len(word2) {
			b.WriteByte(word2[i])
		}
	}
	return b.String()
}
