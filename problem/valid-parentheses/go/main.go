package main

func isValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if open, isClose := pairs[char]; isClose {
			// 閉じ括弧の場合、スタックからポップして対応する開き括弧か確認
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 開き括弧はスタックにプッシュ
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
