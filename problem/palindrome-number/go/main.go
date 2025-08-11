package main

func isPalindrome(x int) bool {
	// 負の数または末尾が0の数は回文にならない
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	left := x
	for left > reversed {
		// 左の末尾一桁を取り出し、逆順にした値の右端に追加
		reversed = reversed*10 + left%10
		// 左の末尾一桁を削除
		left /= 10
	}

	// 偶数桁の場合、左側と右側を比較
	// 奇数桁の場合、元の数字の真ん中(右半分の末尾)は不要なため、一桁削る
	return left == reversed || left == reversed/10
}
