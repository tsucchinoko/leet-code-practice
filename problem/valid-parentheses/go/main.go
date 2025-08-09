package main

func isValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if openPair, isClose := pairs[char]; isClose {
			// 閉じ括弧の場合、スタックの一番上の値が対応する開き括弧と一致か確認
			if len(stack) == 0 || stack[len(stack)-1] != openPair {
				return false
			}
			// popしてスタックを更新
			stack = stack[:len(stack)-1]
		} else {
			// 閉じ括弧以外の文字列はスタックにプッシュ
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
