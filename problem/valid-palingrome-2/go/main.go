package main

func validPalindrome(s string) bool {
	// 部分的に回文か判定する補助関数（インデックスのみ使用）
	isPalindromeRange := func(s string, left, right int) bool {
		for left < right {
			// 違う文字が見つかるまで繰り返す
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}

		// すべて同じ文字
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// 左側を削除した場合 or 右側を削除した場合 のどちらかがOKなら回文
			// そのループのポインタを進行方向に一つ進める=そのループのスキップ=削除したことになる
			return isPalindromeRange(s, left+1, right) ||
				isPalindromeRange(s, left, right-1)
		}
		left++
		right--
	}

	// すべて同じ文字
	return true
}
